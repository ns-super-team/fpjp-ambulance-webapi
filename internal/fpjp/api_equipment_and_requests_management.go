/*
 * Hospital Equipment Management API
 *
 * Equipment and requests management system for hospital
 *
 * API version: 1.0.0
 * Contact: example@mail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package fpjp

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

type EquipmentAndRequestsManagementAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // AddRoomEquipment - Adds new equipment to a room
   AddRoomEquipment(ctx *gin.Context)

    // AddRoomRequest - Adds new request to a room
   AddRoomRequest(ctx *gin.Context)

    // DeleteEquipment - Deletes specific equipment
   DeleteEquipment(ctx *gin.Context)

    // DeleteRequest - Deletes specific request
   DeleteRequest(ctx *gin.Context)

    // GetDepartmentEquipment - Provides list of all equipment in a department
   GetDepartmentEquipment(ctx *gin.Context)

    // GetDepartmentRequests - Provides list of all requests in a department
   GetDepartmentRequests(ctx *gin.Context)

    // GetDepartments - Provides list of all departments
   GetDepartments(ctx *gin.Context)

    // UpdateEquipment - Updates specific equipment
   UpdateEquipment(ctx *gin.Context)

    // UpdateRequest - Updates specific request
   UpdateRequest(ctx *gin.Context)

 }

// partial implementation of EquipmentAndRequestsManagementAPI - all functions must be implemented in add on files
type implEquipmentAndRequestsManagementAPI struct {

}

func newEquipmentAndRequestsManagementAPI() EquipmentAndRequestsManagementAPI {
  return &implEquipmentAndRequestsManagementAPI{}
}

func (this *implEquipmentAndRequestsManagementAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/rooms/:roomId/equipment", this.AddRoomEquipment)
  routerGroup.Handle( http.MethodPost, "/rooms/:roomId/requests", this.AddRoomRequest)
  routerGroup.Handle( http.MethodDelete, "/equipment/:equipmentId", this.DeleteEquipment)
  routerGroup.Handle( http.MethodDelete, "/requests/:requestId", this.DeleteRequest)
  routerGroup.Handle( http.MethodGet, "/departments/:departmentId/equipment", this.GetDepartmentEquipment)
  routerGroup.Handle( http.MethodGet, "/departments/:departmentId/requests", this.GetDepartmentRequests)
  routerGroup.Handle( http.MethodGet, "/departments/", this.GetDepartments)
  routerGroup.Handle( http.MethodPut, "/equipment/:equipmentId", this.UpdateEquipment)
  routerGroup.Handle( http.MethodPut, "/requests/:requestId", this.UpdateRequest)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // AddRoomEquipment - Adds new equipment to a room
// func (this *implEquipmentAndRequestsManagementAPI) AddRoomEquipment(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // AddRoomRequest - Adds new request to a room
// func (this *implEquipmentAndRequestsManagementAPI) AddRoomRequest(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteEquipment - Deletes specific equipment
// func (this *implEquipmentAndRequestsManagementAPI) DeleteEquipment(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteRequest - Deletes specific request
// func (this *implEquipmentAndRequestsManagementAPI) DeleteRequest(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetDepartmentEquipment - Provides list of all equipment in a department
// func (this *implEquipmentAndRequestsManagementAPI) GetDepartmentEquipment(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetDepartmentRequests - Provides list of all requests in a department
// func (this *implEquipmentAndRequestsManagementAPI) GetDepartmentRequests(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetDepartments - Provides list of all departments
// func (this *implEquipmentAndRequestsManagementAPI) GetDepartments(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateEquipment - Updates specific equipment
// func (this *implEquipmentAndRequestsManagementAPI) UpdateEquipment(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateRequest - Updates specific request
// func (this *implEquipmentAndRequestsManagementAPI) UpdateRequest(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
