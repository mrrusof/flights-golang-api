package routing

import (
  "github.com/gin-gonic/gin"
  "github.com/mrrusof/flights-golang-api/controllers"
)

func InitRoutes(server *gin.Engine) {
  server.POST("/flights/summarize", controllers.SummarizeFlight)
}
