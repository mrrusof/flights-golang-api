package main

import (
  "github.com/gin-gonic/gin"
  "github.com/mrrusof/flights-golang-api/routing"
)

func main() {
  server := gin.Default()
  routing.InitRoutes(server)
  server.Run()
}
