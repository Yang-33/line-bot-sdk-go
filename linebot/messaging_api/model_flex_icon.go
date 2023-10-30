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

import (
	"encoding/json"
)

// FlexIcon
// FlexIcon
// https://developers.line.biz/en/reference/messaging-api/#icon
type FlexIcon struct {
	FlexComponent

	/**
	 * Get Url
	 */
	Url string `json:"url"`

	/**
	 * Get Size
	 */
	Size string `json:"size,omitempty"`

	/**
	 * Get AspectRatio
	 */
	AspectRatio string `json:"aspectRatio,omitempty"`

	/**
	 * Get Margin
	 */
	Margin string `json:"margin,omitempty"`

	/**
	 * Get Position
	 */
	Position FlexIconPOSITION `json:"position,omitempty"`

	/**
	 * Get OffsetTop
	 */
	OffsetTop string `json:"offsetTop,omitempty"`

	/**
	 * Get OffsetBottom
	 */
	OffsetBottom string `json:"offsetBottom,omitempty"`

	/**
	 * Get OffsetStart
	 */
	OffsetStart string `json:"offsetStart,omitempty"`

	/**
	 * Get OffsetEnd
	 */
	OffsetEnd string `json:"offsetEnd,omitempty"`

	/**
	 * Get Scaling
	 */
	Scaling bool `json:"scaling"`
}

// MarshalJSON customizes the JSON serialization of the FlexIcon struct.
func (r *FlexIcon) MarshalJSON() ([]byte, error) {

	type Alias FlexIcon
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "icon",
	})
}

// FlexIconPOSITION type

type FlexIconPOSITION string

// FlexIconPOSITION constants
const (
	FlexIconPOSITION_RELATIVE FlexIconPOSITION = "relative"

	FlexIconPOSITION_ABSOLUTE FlexIconPOSITION = "absolute"
)
