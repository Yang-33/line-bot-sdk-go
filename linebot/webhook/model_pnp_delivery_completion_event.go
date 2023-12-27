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
	"fmt"
)

// PnpDeliveryCompletionEvent
// When a request is made to the LINE notification messages API and delivery of the LINE notification message to the user is completed, a dedicated webhook event (delivery completion event) is sent from the LINE Platform to the webhook URL of the bot server.

type PnpDeliveryCompletionEvent struct {
	Event

	/**
	 * Get Source
	 */
	Source SourceInterface `json:"source,omitempty"`

	/**
	 * Time of the event in milliseconds. (Required)
	 */
	Timestamp int64 `json:"timestamp"`

	/**
	 * Get Mode
	 */
	Mode EventMode `json:"mode"`

	/**
	 * Webhook Event ID. An ID that uniquely identifies a webhook event. This is a string in ULID format. (Required)
	 */
	WebhookEventId string `json:"webhookEventId"`

	/**
	 * Get DeliveryContext
	 */
	DeliveryContext *DeliveryContext `json:"deliveryContext"`

	/**
	 * Get Delivery
	 */
	Delivery *PnpDelivery `json:"delivery"`
}

func (cr *PnpDeliveryCompletionEvent) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return fmt.Errorf("JSON parse error in map: %w", err)
	}

	if raw["type"] != nil {

		err = json.Unmarshal(raw["type"], &cr.Type)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(Type): %w", err)
		}

	}

	if raw["source"] != nil {

		if rawsource, ok := raw["source"]; ok && rawsource != nil {
			Source, err := UnmarshalSource(rawsource)
			if err != nil {
				return fmt.Errorf("JSON parse error in Source(discriminator): %w", err)
			}
			cr.Source = Source
		}

	}

	if raw["timestamp"] != nil {

		err = json.Unmarshal(raw["timestamp"], &cr.Timestamp)
		if err != nil {
			return fmt.Errorf("JSON parse error in int64(Timestamp): %w", err)
		}

	}

	if raw["mode"] != nil {

		err = json.Unmarshal(raw["mode"], &cr.Mode)
		if err != nil {
			return fmt.Errorf("JSON parse error in EventMode(Mode): %w", err)
		}

	}

	if raw["webhookEventId"] != nil {

		err = json.Unmarshal(raw["webhookEventId"], &cr.WebhookEventId)
		if err != nil {
			return fmt.Errorf("JSON parse error in string(WebhookEventId): %w", err)
		}

	}

	if raw["deliveryContext"] != nil {

		err = json.Unmarshal(raw["deliveryContext"], &cr.DeliveryContext)
		if err != nil {
			return fmt.Errorf("JSON parse error in DeliveryContext(DeliveryContext): %w", err)
		}

	}

	if raw["delivery"] != nil {

		err = json.Unmarshal(raw["delivery"], &cr.Delivery)
		if err != nil {
			return fmt.Errorf("JSON parse error in PnpDelivery(Delivery): %w", err)
		}

	}

	return nil
}

// MarshalJSON customizes the JSON serialization of the PnpDeliveryCompletionEvent struct.
func (r *PnpDeliveryCompletionEvent) MarshalJSON() ([]byte, error) {

	r.Source = setDiscriminatorPropertySource(r.Source)

	type Alias PnpDeliveryCompletionEvent
	return json.Marshal(&struct {
		*Alias

		Type string `json:"type"`
	}{
		Alias: (*Alias)(r),

		Type: "delivery",
	})
}