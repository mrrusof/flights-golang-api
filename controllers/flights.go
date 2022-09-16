package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/mrrusof/flights-golang-api/services"
  "net/http"
)

func SummarizeFlight(c *gin.Context) {
  var flight services.FlightPath

  if err := c.BindJSON(&flight); err != nil {
    c.Writer.WriteHeader(http.StatusBadRequest)
    return
  }

  if len(flight.Path) == 0 {
    c.Writer.WriteHeader(http.StatusBadRequest)
    return
  }

  c.JSON(http.StatusOK, services.Summarize(flight))
}
