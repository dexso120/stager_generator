package handlers

import (
	"os"
	"log"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"stager_generator/jsobfuscation1"
	"stager_generator/jsobfuscation2"
)

const (
	OUTFILE_PATH = "./outfile"
)

type Action struct {
    Id                  int `json:"id"`
    CommandString       string `json:"command"`
    Outfile             string `json:"outfile"`
}

func CallObfuscation(c *gin.Context){
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

	outfilePath := OUTFILE_PATH + "/" + d.Outfile

	err := os.WriteFile(outfilePath, []byte(obfuscationResultString), 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result written to file:", outfilePath)

	// Returning output file as attachment
	c.FileAttachment(outfilePath, d.Outfile)
}