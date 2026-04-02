package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context){
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
	  "message": "pong",
	})
}

func GetHome(c *gin.Context){
	c.HTML(http.StatusOK, "pages/home.html", gin.H{
		"Active": "home",
		"Title": "Stager Generator App",
		"Message": "Welcome to my app.",
	})
}

type ObfuscationType struct {
    ID    int
    Label string
}

func GetObfuscationPage(c *gin.Context){
	c.HTML(http.StatusOK, "pages/obfuscation_page.html", gin.H{
		"Active": "obfuscation_page",
		"Types": []ObfuscationType{
			{ID: 1, Label: "JScript Obfuscation 1"},
			{ID: 2, Label: "JScript Obfuscation 2"},
		},
	})
}