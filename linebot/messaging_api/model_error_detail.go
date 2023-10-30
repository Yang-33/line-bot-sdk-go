/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package messaging_api

// ErrorDetail
// ErrorDetail

type ErrorDetail struct {

	/**
	 * Details of the error. Not included in the response under certain situations.
	 */
	Message string `json:"message,omitempty"`

	/**
	 * Location of where the error occurred. Returns the JSON field name or query parameter name of the request. Not included in the response under certain situations.
	 */
	Property string `json:"property,omitempty"`
}
