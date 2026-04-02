package main

import (
  "log"

  "github.com/gin-gonic/gin"

  "stager_generator/routes"
)

func main() {
  // Gin Routes
  r := gin.Default()
  routes.Register(r)
  
  if err := r.Run(":8888"); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}