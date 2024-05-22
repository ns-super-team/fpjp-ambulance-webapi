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

type Request struct {

	// Unique identifier of the request
	Id string `json:"id"`

	// Identifier of the room the request is associated with
	RoomId string `json:"room_id"`

	// Type of the request
	Type string `json:"type"`

	// Name of the equipment requested or to be repaired
	Name string `json:"name"`

	// Number of items requested (only applicable for missing-equipment requests)
	Count *int32 `json:"count,omitempty"`

	// Detailed description of the request
	Description string `json:"description"`
}
