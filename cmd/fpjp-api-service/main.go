package main

import (
    "log"
    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/ns-super-team/fpjp-ambulance-webapi/api"
		"github.com/ns-super-team/fpjp-ambulance-webapi/internal/fpjp"
)

func main() {
    log.Printf("Server started")
    port := os.Getenv("FPJP_API_PORT")
    if port == "" {
        port = "8080"
    }
    environment := os.Getenv("FPJP_API_ENVIRONMENT")
    if !strings.EqualFold(environment, "production") { // case insensitive comparison
        gin.SetMode(gin.DebugMode)
    }
    engine := gin.New()
    engine.Use(gin.Recovery())
    // request routings
		fpjp.AddRoutes(engine)
    engine.GET("/openapi", api.HandleOpenApi)
    engine.Run(":" + port)
}