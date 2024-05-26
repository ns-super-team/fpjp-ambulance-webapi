package fpjp

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ns-super-team/fpjp-ambulance-webapi/internal/db_service"
	"go.mongodb.org/mongo-driver/bson"
)

// AddRoomEquipment - Adds new equipment to a room
func (this *implEquipmentAndRequestsManagementAPI) AddRoomEquipment(ctx *gin.Context) {
	fmt.Println("req -> AddRoomEquipment")

	value, exists := ctx.Get("equipment_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "equipment_service not found",
				"error":   "equipment_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Equipment])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "equipment_service context is not of required type",
				"error":   "cannot cast equipment_service context to db_service.DbService",
			})
		return
	}

	equipment := Equipment{}
	err := ctx.ShouldBindJSON(&equipment)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	// get room ID from URL param
	URLroomId := ctx.Param("roomId")

	// check if room ID from URL param and room ID from request body are equal
	if URLroomId != equipment.RoomId {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Room ID provided in request body is not equal to room ID in URL parameter.",
				"error":   "Room ID provided in request body is not equal to room ID in URL parameter.",
			})
		return
	}

	// create new UUID
	if equipment.Id == "" {
		equipment.Id = uuid.New().String()
	}

	// create equipment
	err = db.CreateDocument(ctx, equipment.Id, &equipment)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			equipment,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Equipment already exists",
				"error":   err.Error(),
			})
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create equipment in database",
				"error":   err.Error(),
			})
	}
}

// DeleteEquipment - Deletes specific equipment
func (this *implEquipmentAndRequestsManagementAPI) DeleteEquipment(ctx *gin.Context) {
	fmt.Println("req -> DeleteEquipment")

	value, exists := ctx.Get("equipment_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "equipment_service not found",
				"error":   "equipment_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Equipment])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "equipment_service context is not of type db_service.DbService",
				"error":   "cannot cast equipment_service context to db_service.DbService",
			})
		return
	}

	// get equipment ID from URL
	equipmentId := ctx.Param("equipmentId")

	// delete to document
	err := db.DeleteDocument(ctx, equipmentId)

	switch err {
	case nil:
		ctx.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Equipment not found",
				"error":   err.Error(),
			})
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete equipment from database",
				"error":   err.Error(),
			})
	}
}

// UpdateEquipment - Updates specific equipment
func (this *implEquipmentAndRequestsManagementAPI) UpdateEquipment(ctx *gin.Context) {
	fmt.Println("req -> UpdateEquipment")

	value, exists := ctx.Get("equipment_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "equipment_service not found",
				"error":   "equipment_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Equipment])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "equipment_service context is not of type db_service.DbService",
				"error":   "cannot cast equipment_service context to db_service.DbService",
			})
		return
	}

	equipment := Equipment{}
	err := ctx.ShouldBindJSON(&equipment)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	// get equipment ID from URL param
	URLequipmentId := ctx.Param("equipmentId")

	// check if ID from URL param and ID from request body are equal
	if URLequipmentId != equipment.Id {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "ID provided in request body is not equal to ID in URL parameter.",
				"error":   "ID provided in request body is not equal to ID in URL parameter.",
			})
		return
	}

	// update equipment
	err = db.UpdateDocument(ctx, equipment.Id, &equipment)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			equipment,
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Equipment with provided ID was not found.",
				"error":   err.Error(),
			})
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update equipment in database.",
				"error":   err.Error(),
			})
	}
}

// AddRoomRequest - Adds new request to a room
func (this *implEquipmentAndRequestsManagementAPI) AddRoomRequest(ctx *gin.Context) {
	fmt.Println("req -> AddRoomRequest")

	value, exists := ctx.Get("request_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "request_service not found",
				"error":   "request_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Request])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "request_service context is not of required type",
				"error":   "cannot cast request_service context to db_service.DbService",
			})
		return
	}

	request := Request{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	// Get room ID from URL param
	URLroomId := ctx.Param("roomId")

	// Check if room ID from URL param and room ID from request body are equal
	if URLroomId != request.RoomId {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Room ID provided in request body is not equal to room ID in URL parameter.",
				"error":   "Room ID provided in request body is not equal to room ID in URL parameter.",
			})
		return
	}

	// Create new UUID if request Id is empty
	if request.Id == "" {
		request.Id = uuid.New().String()
	}

	// Create request
	err = db.CreateDocument(ctx, request.Id, &request)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			request,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Request already exists",
				"error":   err.Error(),
			})
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create request in database",
				"error":   err.Error(),
			})
	}
}

// DeleteRequest - Deletes specific request
func (this *implEquipmentAndRequestsManagementAPI) DeleteRequest(ctx *gin.Context) {
	fmt.Println("req -> DeleteRequest")

	value, exists := ctx.Get("request_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "request_service not found",
				"error":   "request_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Request])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "request_service context is not of type db_service.DbService",
				"error":   "cannot cast request_service context to db_service.DbService",
			})
		return
	}

	// Get request ID from URL
	requestId := ctx.Param("requestId")

	// Delete the document
	err := db.DeleteDocument(ctx, requestId)

	switch err {
	case nil:
		ctx.AbortWithStatus(http.StatusNoContent)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Request not found",
				"error":   err.Error(),
			})
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete request from database",
				"error":   err.Error(),
			})
	}
}

// UpdateRequest - Updates specific request
func (this *implEquipmentAndRequestsManagementAPI) UpdateRequest(ctx *gin.Context) {
	fmt.Println("req -> UpdateRequest")

	value, exists := ctx.Get("request_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "request_service not found",
				"error":   "request_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Request])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "request_service context is not of type db_service.DbService",
				"error":   "cannot cast request_service context to db_service.DbService",
			})
		return
	}

	request := Request{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	// Get request ID from URL param
	URLrequestId := ctx.Param("requestId")

	// Check if ID from URL param and ID from request body are equal
	if URLrequestId != request.Id {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "ID provided in request body is not equal to ID in URL parameter.",
				"error":   "ID provided in request body is not equal to ID in URL parameter.",
			})
		return
	}

	// Update request
	err = db.UpdateDocument(ctx, request.Id, &request)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			request,
		)
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Request with provided ID was not found.",
				"error":   err.Error(),
			})
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update request in database.",
				"error":   err.Error(),
			})
	}
}

// GetDepartments - Provides list of all departments
func (this *implEquipmentAndRequestsManagementAPI) GetDepartments(ctx *gin.Context) {
	fmt.Println("req -> GetDepartments")

	value, exists := ctx.Get("department_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "department_service not found",
				"error":   "department_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Department])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "department_service context is not of type db_service.DbService",
				"error":   "cannot cast department_service context to db_service.DbService",
			})
		return
	}

	//create empty filter
	filter := bson.M{}

	// get all departments
	departments, err := db.FindDocuments(ctx, filter)

	if len(departments) == 0 {
		fmt.Printf("empty")
	}

	switch err {
	case nil:
		ctx.JSON(
			http.StatusOK,
			departments,
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to get departments",
				"error":   err.Error(),
			})
	}
}

// GetDepartmentEquipment - Provides list of all equipment in a department
func (this *implEquipmentAndRequestsManagementAPI) GetDepartmentEquipment(ctx *gin.Context) {
	fmt.Println("req -> GetDepartmentEquipment")

	// get department ID from URL parameter
	departmentID := ctx.Param("departmentId")
	if departmentID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Department ID is required",
			"error":   "Department ID is missing in URL parameters",
		})
		return
	}

	// department service
	value, exists := ctx.Get("department_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "department_service not found",
				"error":   "department_service not found",
			})
		return
	}

	departmentService, ok := value.(db_service.DbService[Department])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "department_service context is not of type db_service.DbService",
				"error":   "cannot cast department_service context to db_service.DbService",
			})
		return
	}

	// get department
	department, err := departmentService.FindDocument(ctx, departmentID)

	switch err {
	case nil:
		// do nothing
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Department with provided ID was not found.",
				"error":   err.Error(),
			})
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to find Department in database.",
				"error":   err.Error(),
			})
	}

	// room service
	value, exists = ctx.Get("room_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "room_service not found",
				"error":   "room_service not found",
			})
		return
	}

	roomService, ok := value.(db_service.DbService[Room])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "room_service context is not of type db_service.DbService",
				"error":   "cannot cast room_service context to db_service.DbService",
			})
		return
	}

	// equipment service
	value, exists = ctx.Get("equipment_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "equipment_service not found",
				"error":   "equipment_service not found",
			})
		return
	}

	equipmentService, ok := value.(db_service.DbService[Equipment])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "equipment_service context is not of type db_service.DbService",
				"error":   "cannot cast equipment_service context to db_service.DbService",
			})
		return
	}

	// filter rooms by department
	roomsFilter := bson.M{"department_id": departmentID}

	// get rooms
	rooms, err := roomService.FindDocuments(ctx, roomsFilter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Failed to retrieve rooms",
			"error":   err.Error(),
		})
		return
	}

	// get list of room IDs
	roomIDs := make([]string, len(rooms))
	for i, room := range rooms {
		roomIDs[i] = room.Id
	}

	// get equipment based on room IDs
	equipmentFilter := bson.M{"room": bson.M{"$in": roomIDs}}
	equipment, err := equipmentService.FindDocuments(ctx, equipmentFilter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Failed to retrieve equipment",
			"error":   err.Error(),
		})
		return
	}

	// create response object
	response := struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Rooms []struct {
			Id        string      `json:"id"`
			Name      string      `json:"name"`
			Equipment []Equipment `json:"equipment"`
		} `json:"rooms"`
	}{
		Id:   departmentID,
		Name: department.Name,
	}

	for _, room := range rooms {
		roomEquip := []Equipment{}
		for _, eq := range equipment {
			fmt.Println("eq.RoomId", eq.RoomId)
			fmt.Println("room.Id", room.Id)
			fmt.Println(eq.RoomId == room.Id)
			if eq.RoomId == room.Id {
				roomEquip = append(roomEquip, *eq)
			}
		}
		response.Rooms = append(response.Rooms, struct {
			Id        string      `json:"id"`
			Name      string      `json:"name"`
			Equipment []Equipment `json:"equipment"`
		}{
			Id:        room.Id,
			Name:      room.Name,
			Equipment: roomEquip,
		})
	}
	fmt.Println(response)

	ctx.JSON(http.StatusOK, response)
}

// GetDepartmentRequests - Provides list of all requests in a department
func (this *implEquipmentAndRequestsManagementAPI) GetDepartmentRequests(ctx *gin.Context) {
	fmt.Println("req -> GetDepartmentRequests")

	departmentID := ctx.Param("departmentId")
	if departmentID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "Department ID is required",
			"error":   "Department ID is missing in URL parameters",
		})
		return
	}

	// department service
	value, exists := ctx.Get("department_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "department_service not found",
				"error":   "department_service not found",
			})
		return
	}

	departmentService, ok := value.(db_service.DbService[Department])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "department_service context is not of type db_service.DbService",
				"error":   "cannot cast department_service context to db_service.DbService",
			})
		return
	}

	// get department
	department, err := departmentService.FindDocument(ctx, departmentID)

	// room service
	value, exists = ctx.Get("room_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "room_service not found",
				"error":   "room_service not found",
			})
		return
	}

	roomService, ok := value.(db_service.DbService[Room])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "room_service context is not of type db_service.DbService",
				"error":   "cannot cast room_service context to db_service.DbService",
			})
		return
	}

	// request service
	value, exists = ctx.Get("request_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "request_service not found",
				"error":   "request_service not found",
			})
		return
	}

	requestService, ok := value.(db_service.DbService[Request])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "request_service context is not of type db_service.DbService",
				"error":   "cannot cast request_service context to db_service.DbService",
			})
		return
	}

	// filter rooms by department
	roomsFilter := bson.M{"department_id": departmentID}

	// get rooms
	rooms, err := roomService.FindDocuments(ctx, roomsFilter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Failed to retrieve rooms",
			"error":   err.Error(),
		})
		return
	}

	// get list of room IDs
	roomIDs := make([]string, len(rooms))
	for i, room := range rooms {
		roomIDs[i] = room.Id
	}

	// get requests based on room IDs
	requestFilter := bson.M{"room": bson.M{"$in": roomIDs}}
	requests, err := requestService.FindDocuments(ctx, requestFilter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Internal Server Error",
			"message": "Failed to retrieve requests",
			"error":   err.Error(),
		})
		return
	}

	// create response object
	response := struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Rooms []struct {
			Id       string    `json:"id"`
			Name     string    `json:"name"`
			Requests []Request `json:"requests"`
		} `json:"rooms"`
	}{
		Id:   departmentID,
		Name: department.Name,
	}

	for _, room := range rooms {
		roomRequests := []Request{}
		for _, req := range requests {
			if req.RoomId == room.Id {
				roomRequests = append(roomRequests, *req)
			}
		}
		response.Rooms = append(response.Rooms, struct {
			Id       string    `json:"id"`
			Name     string    `json:"name"`
			Requests []Request `json:"requests"`
		}{
			Id:       room.Id,
			Name:     room.Name,
			Requests: roomRequests,
		})
	}

	ctx.JSON(http.StatusOK, response)
}
