package routes

import (
	"github.com/gin-gonic/gin"

	"stager_generator/handlers"
)

func Register(r *gin.Engine) {

	// Loading HTML template
	r.LoadHTMLGlob("templates/**/*")

	// Simple ping
	r.GET("/ping", handlers.Ping)

	// Home page
	r.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/home")
    })

	r.GET("/home", handlers.GetHome)

	// Maing page for obfuscations
	r.GET("/obfuscation_page", handlers.GetObfuscationPage)

	// Obfuscation APIs
	r.POST("/obfuscation", handlers.CallObfuscation)

}