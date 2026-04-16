package main

import (
  "log"
  "path/filepath"

  "github.com/gin-gonic/gin"

  "stager_generator/routes"
  "stager_generator/utils"
)

func main() {

  // Creating outfile and temp_outfile directory
  if err := utils.EnsureDir(filepath.Join("./", "outfile")); err != nil {
    log.Fatal(err)
  }

  if err := utils.EnsureDir(filepath.Join("./", "temp_outfile")); err != nil {
    log.Fatal(err)
  }

  // Gin Routes
  r := gin.Default()
  routes.Register(r)
  
  if err := r.Run(":8888"); err != nil {
    log.Fatalf("failed to run server: %v", err)
  }
}