/**
 * Webhook Type Definition
 * Webhook event definition of the LINE Messaging API
 *
 * The version of the OpenAPI document: 1.0.0
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
package webhook

import (
	"encoding/json"
)

// ScenarioResultThingsContent
// ScenarioResultThingsContent

type ScenarioResultThingsContent struct {
	ThingsContent

	/**
	 * Device ID of the device that has been linked with LINE. (Required)
	 */
	DeviceId string `json:"deviceId"`

	/**
	 * Get Result
	 */
	Result *ScenarioResult `json:"result"`
}

// MarshalJSON customizes the JSON serialization of the ScenarioResultThingsContent struct.
func (r *ScenarioResultThingsContent) MarshalJSON() ([]byte, error) {

	type Alias ScenarioResultThingsContent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type,omitempty"`
	}{
		Alias: (*Alias)(r),

		Type: "scenarioResult",
	})
}