package handlers

import (
	"os"
	"log"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"stager_generator/utils"
	jsobfuscation1 "stager_generator/obfuscation/jsobfuscation1"
	jsobfuscation2 "stager_generator/obfuscation/jsobfuscation2"
	psobfuscation1 "stager_generator/obfuscation/psobfuscation1"
)

const (
	OUTFILE_PATH = "./outfile"
)

type Action struct {
    Id                  int `json:"id"`
    CommandString       string `json:"command"`
    Outfile             string `json:"outfile"`
}

func CallJScriptObfuscation(c *gin.Context){
	// Clean up previous output files
	utils.CleanUpOutfiles()

	var d Action
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
	}

	var obfuscationResultString string

	switch d.Id {
	case 1:
		obfuscationResultString = jsobfuscation1.JsObfuscation1(d.CommandString)
	case 2:
		obfuscationResultString = jsobfuscation2.JsObfuscation2(d.CommandString)
	default:
		obfuscationResultString = jsobfuscation1.JsObfuscation1(d.CommandString)
	}

	outfilePath := filepath.Join(OUTFILE_PATH, "/" + d.Outfile)

	err := os.WriteFile(outfilePath, []byte(obfuscationResultString), 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result written to file:", outfilePath)

	// Returning output file as attachment
	c.FileAttachment(outfilePath, d.Outfile)
}

func CallPowershellObfuscation(c *gin.Context){

	// Clean up previous output files
	utils.CleanUpOutfiles()

	// Input fields based on chosen obfuscation type
	// TODO: add ID field
	textFields  := []string{"localFilepath", "url", "getType", "getMethod", "args"}
    fileFields  := []string{"uploadFile", "payloadFile"}

    parsed := gin.H{}

    // Collect all text fields that were submitted
    for _, field := range textFields {
        if val := c.PostForm(field); val != "" {
            parsed[field] = val
        }
    }

    // Collect all file fields that were submitted
    for _, field := range fileFields {
        file, err := c.FormFile(field)
        if err == nil && file != nil {
            dst := filepath.Join("./temp_outfile", file.Filename)
            if err := c.SaveUploadedFile(file, dst); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save " + field})
                return
            }
            //parsed[field] = file.Filename
            parsed[field] = dst
        }
    }

    // Calling Powershell Obfuscation Function
    fmt.Println("Test local file path", parsed["localFilepath"])
    var obfuscationResultString string

    // Assert form data types
    uploadDotnetFilePath := parsed["uploadFile"].(string)
    localFilePath := parsed["localFilepath"].(string)
    url := parsed["url"].(string)
    getType := parsed["getType"].(string)
    getMethod := parsed["getMethod"].(string)

    var args string
    if parsed["args"] != nil{
    	args = parsed["args"].(string)
    }else{
    	args = ""
    }

    obfuscationResultString, tempDotnetFilepath := psobfuscation1.PsObfuscation1(uploadDotnetFilePath, localFilePath, url, getType, getMethod, args)

    /*

	switch d.Id {
	case 1:
		obfuscationResultString = psobfuscation1.PsObfuscation1(parsed["uploadDotnetFile"], parsed["localFilepath"], parsed["url"], parsed["getType"], parsed["getMethod"])
	default:
		obfuscationResultString = psobfuscation1.PsObfuscation1(parsed["uploadDotnetFile"], parsed["localFilepath"], parsed["url"], parsed["getType"], parsed["getMethod"])
	}
	*/

	loaderOutfilePath := filepath.Join(OUTFILE_PATH, "/script.ps1")

	err := os.WriteFile(loaderOutfilePath, []byte(obfuscationResultString), 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result loader written to file:", loaderOutfilePath)

	// Zipping the loader script and the encoded .NET executable
	zipFilePath := filepath.Join(OUTFILE_PATH, "/output.zip")

	utils.Zipfiles(zipFilePath, loaderOutfilePath, tempDotnetFilepath)

	utils.CleanUpTempOutfiles()

	// Returning output file as attachment
	c.FileAttachment(zipFilePath, "output.zip")
}