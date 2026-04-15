package routes

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"stager_generator/handlers"
)

func Register(r *gin.Engine) {

	// Loading HTML template
	r.LoadHTMLGlob("templates/**/*")

	// Simple ping
	r.GET("/ping", handlers.Ping)

	// GUI Pages //

	// Home page
	r.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/home")
    })

	r.GET("/home", handlers.GetHome)

	// Maing page for obfuscations
	r.GET("/obfuscation_page", handlers.GetObfuscationPage)

	r.GET("/ps_loader_page", handlers.GetPsLoaderPage)

	// GUI Pages End //

	// Backend API Endpoints //
	
	// Obfuscation APIs
	r.POST("/jscript_obfuscation", handlers.CallJScriptObfuscation)

	r.POST("/powershell_obfuscation", handlers.CallPowershellObfuscation)

	// Backend API Endpoints End //

}