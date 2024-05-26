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

type Equipment struct {

	// Unique identifier of the equipment
	Id string `json:"id" bson:"id"`

	// Identifier of the room the equipment belongs to
	Room string `json:"room" bson:"room" binding:"required"`

	// Type of the equipment
	Type string `json:"type" bson:"type" binding:"required"`

	// Name of the equipment
	Name string `json:"name" bson:"name" binding:"required"`

	// Number of equipment items available
	Count int32 `json:"count" bson:"count" binding:"required"`
}
