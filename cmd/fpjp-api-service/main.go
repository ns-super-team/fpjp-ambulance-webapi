package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ns-super-team/fpjp-ambulance-webapi/api"
	"github.com/ns-super-team/fpjp-ambulance-webapi/internal/db_service"
	"github.com/ns-super-team/fpjp-ambulance-webapi/internal/fpjp"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("AMBULANCE_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("AMBULANCE_API_ENVIRONMENT")
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
		MaxAge:           12 * time.Hour,
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

	// db initialization
	if environment == "development" {
		insertInitialData(departmentService, roomService)
	}

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

// populates Departments and Rooms with initial data, if these collections are empty
func insertInitialData(departmentService db_service.DbService[fpjp.Department], roomService db_service.DbService[fpjp.Room]) {
	ctx := context.Background()

	departments, err := departmentService.FindDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Error checking departments collection: %v", err)
	}
	if len(departments) > 0 {
		return
	}

	rooms, err := roomService.FindDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Error checking rooms collection: %v", err)
	}
	if len(rooms) > 0 {
		return
	}

	initialDepartments := []fpjp.Department{
		{Id: "1", Name: "Pediatrické oddelenie"},
		{Id: "2", Name: "Chirurgia"},
		{Id: "3", Name: "Alergológia"},
		{Id: "4", Name: "Ortopédia"},
		{Id: "5", Name: "Neurológia"},
	}

	initialRooms := []fpjp.Room{
		{Id: "1", DepartmentId: "1", Name: "Miestnosť 1.1"},
		{Id: "2", DepartmentId: "1", Name: "Miestnosť 1.2"},
		{Id: "3", DepartmentId: "2", Name: "Miestnosť 2.1"},
		{Id: "4", DepartmentId: "2", Name: "Miestnosť 2.2"},
		{Id: "5", DepartmentId: "2", Name: "Miestnosť 2.3"},
		{Id: "6", DepartmentId: "3", Name: "Miestnosť 3.1"},
		{Id: "7", DepartmentId: "4", Name: "Miestnosť 4.1"},
		{Id: "8", DepartmentId: "5", Name: "Miestnosť 5.1"},
		{Id: "9", DepartmentId: "5", Name: "Miestnosť 5.2"},
	}

	for _, department := range initialDepartments {
		err := departmentService.CreateDocument(ctx, department.Id, &department)
		if err != nil {
			log.Fatalf("Failed to insert department %s: %v", department.Id, err)
		}
	}

	for _, room := range initialRooms {
		err := roomService.CreateDocument(ctx, room.Id, &room)
		if err != nil {
			log.Fatalf("Failed to insert room %s: %v", room.Id, err)
		}
	}
}
