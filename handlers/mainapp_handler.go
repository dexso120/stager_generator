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
    Description string
}

func GetObfuscationPage(c *gin.Context){
	c.HTML(http.StatusOK, "pages/obfuscation_page.html", gin.H{
		"Active": "obfuscation_page",
		"Types": []ObfuscationType{
			{
				ID: 1,
				Label: "JScript Obfuscation 1",
				Description: `Description:<br>

				- Inserts comamnd into "template.js" JScript file.<br>
				- "template.js" content is embedded into loader.js as comments (Prefix: "////")<br>
				- "loader.js" will find comments start with "////", regroup it and execute as JScript.`,
			},
			{
				ID: 2,
				Label: "JScript Obfuscation 2",
				Description: `Sample Reference:<br>
				<a href="https://www.swisspost-cybersecurity.ch/news/purelogs-infostealer-analysis-dont-judge-a-png-by-its-header">https://www.swisspost-cybersecurity.ch/news/purelogs-infostealer-analysis-dont-judge-a-png-by-its-header</a><br><br>

				Description:<br>
				- Inserts command into "template.ps1" powershell file<br>
				- Encode powershell into base64 (UTF-16)<br>
				- Insert junk content into encoded powershell<br>
				- Place obfuscated base64 powershell into "loader.js" along with the junk to be removed at runtime.<br>

				`,
			},
		},
	})
}