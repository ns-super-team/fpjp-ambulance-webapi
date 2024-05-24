package fpjp

import (
	"net/http"

	"github.com/ns-super-team/fpjp-ambulance-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


// AddRoomEquipment - Adds new equipment to a room
func (this *implEquipmentAndRequestsManagementAPI) AddRoomEquipment(ctx *gin.Context) {
	value, exists := ctx.Get("equipment_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status": "Internal Server Error",
				"message": "equipment_service not found",
				"error": "equipment_service not found",
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
	err := ctx.ShouldBindJSON(&equipment);
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": "Bad Request",
				"message": "Invalid request body",
				"error": err.Error(),
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
				"status": "Bad Request",
				"message": "Room ID provided in request body is not equal to room ID in URL parameter.",
				"error": "Room ID provided in request body is not equal to room ID in URL parameter.",
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
	value, exists := ctx.Get("equipment_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status": "Internal Server Error",
				"message": "equipment_service not found",
				"error": "equipment_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Equipment])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status": "Internal Server Error",
				"message": "equipment_service context is not of type db_service.DbService",
				"error": "cannot cast equipment_service context to db_service.DbService",
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
				"status": "Not Found",
				"message": "Equipment not found",
				"error": err.Error(),
			})
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status": "Bad Gateway",
				"message": "Failed to delete equipment from database",
				"error": err.Error(),
			})
	}
}


// UpdateEquipment - Updates specific equipment
func (this *implEquipmentAndRequestsManagementAPI) UpdateEquipment(ctx *gin.Context) {
	value, exists := ctx.Get("equipment_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status": "Internal Server Error",
				"message": "equipment_service not found",
				"error": "equipment_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[Equipment])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status": "Internal Server Error",
				"message": "equipment_service context is not of type db_service.DbService",
				"error": "cannot cast equipment_service context to db_service.DbService",
			})
		return
	}

	equipment := Equipment{}
	err := ctx.ShouldBindJSON(&equipment);
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status": "Bad Request",
				"message": "Invalid request body",
				"error": err.Error(),
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
				"status": "Bad Request",
				"message": "ID provided in request body is not equal to ID in URL parameter.",
				"error": "ID provided in request body is not equal to ID in URL parameter.",
			},
		)
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
				"status": "Not Found",
				"message": "Equipment with provided ID was not found.",
				"error": err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status": "Bad Gateway",
				"message": "Failed to update equipment in database.",
				"error": err.Error(),
			},
		)
	}
}


// AddRoomRequest - Adds new request to a room
func (this *implEquipmentAndRequestsManagementAPI) AddRoomRequest(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// DeleteRequest - Deletes specific request
func (this *implEquipmentAndRequestsManagementAPI) DeleteRequest(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// UpdateRequest - Updates specific request
func (this *implEquipmentAndRequestsManagementAPI) UpdateRequest(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetDepartments - Provides list of all departments
func (this *implEquipmentAndRequestsManagementAPI) GetDepartments(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetDepartmentEquipment - Provides list of all equipment in a department
func (this *implEquipmentAndRequestsManagementAPI) GetDepartmentEquipment(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// GetDepartmentRequests - Provides list of all requests in a department
func (this *implEquipmentAndRequestsManagementAPI) GetDepartmentRequests(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusNotImplemented)
}