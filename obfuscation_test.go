package main

import (
	"fmt"
	"testing"
	"log"
	"os"

	"stager_generator/jsobfuscation1"
	"stager_generator/jsobfuscation2"
)

func TestObfuscation1(t *testing.T){
	fmt.Println("Starting TestObfuscation1.")


	testCommand := "cmd.exe /c calc.exe"
	testOutfile := "test_obfuscation_1.js"

	obfuscationResultString := jsobfuscation1.JsObfuscation1(testCommand)

	err := os.WriteFile(testOutfile, []byte(obfuscationResultString), 0644)

    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("Test function end.")
}

func TestObfuscation2(t *testing.T){
	fmt.Println("Starting TestObfuscation2.")


	testCommand := "cmd.exe /c calc.exe"
	testOutfile := "test_obfuscation_2.js"

	obfuscationResultString := jsobfuscation2.JsObfuscation2(testCommand)

	err := os.WriteFile(testOutfile, []byte(obfuscationResultString), 0644)

    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("Test function end.")
}