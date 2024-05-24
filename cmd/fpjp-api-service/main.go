package main

import (
    "log"
    "os"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/ns-super-team/fpjp-ambulance-webapi/api"
		"github.com/ns-super-team/fpjp-ambulance-webapi/internal/fpjp"
		"github.com/ns-super-team/fpjp-ambulance-webapi/internal/db_service"
    "context"
    "time"
    "github.com/gin-contrib/cors"
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
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		MaxAge: 12 * time.Hour,
	})
	engine.Use(corsMiddleware)

	// setup contexts
	departmentService := db_service.NewMongoService[fpjp.Department](db_service.MongoServiceConfig{
		Collection: "departments",
	})
	defer departmentService.Disconnect(context.Background())

	equipmentService := db_service.NewMongoService[fpjp.Equipment](db_service.MongoServiceConfig{
		Collection: "equipment",
	})
	defer equipmentService.Disconnect(context.Background())

	requestService := db_service.NewMongoService[fpjp.Request](db_service.MongoServiceConfig{
		Collection: "requests",
	})
	defer requestService.Disconnect(context.Background())

	roomService := db_service.NewMongoService[fpjp.Room](db_service.MongoServiceConfig{
		Collection: "rooms",
	})
	defer roomService.Disconnect(context.Background())

	// update middleware
	engine.Use(func(ctx *gin.Context) {
		ctx.Set("department_service", departmentService)
		ctx.Set("equipment_service", equipmentService)
		ctx.Set("request_service", requestService)
		ctx.Set("room_service", roomService)
		ctx.Next()
})

	// request routings
	fpjp.AddRoutes(engine)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}