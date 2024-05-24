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

type Department struct {

	// Unique identifier of the department
	Id string `json:"id" bson:"id"`

	// Name of the department
	Name string `json:"name" bson:"name"`
}
