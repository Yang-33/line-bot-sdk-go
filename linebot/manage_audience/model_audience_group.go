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
package manage_audience

// AudienceGroup
// Audience group

type AudienceGroup struct {

	/**
	 * The audience ID.
	 */
	AudienceGroupId int64 `json:"audienceGroupId"`

	/**
	 * Get Type
	 */
	Type AudienceGroupType `json:"type,omitempty"`

	/**
	 * The audience&#39;s name.
	 */
	Description string `json:"description,omitempty"`

	/**
	 * Get Status
	 */
	Status AudienceGroupStatus `json:"status,omitempty"`

	/**
	 * Get FailedType
	 */
	FailedType AudienceGroupFailedType `json:"failedType,omitempty"`

	/**
	 * The number of users included in the audience.
	 */
	AudienceCount int64 `json:"audienceCount"`

	/**
	 * When the audience was created (in UNIX time).
	 */
	Created int64 `json:"created"`

	/**
	 * The request ID that was specified when the audience was created. This is only included when `audienceGroup.type` is CLICK or IMP.
	 */
	RequestId string `json:"requestId,omitempty"`

	/**
	 * The URL that was specified when the audience was created. This is only included when `audienceGroup.type` is CLICK and link URL is specified.
	 */
	ClickUrl string `json:"clickUrl,omitempty"`

	/**
	 * The value indicating the type of account to be sent, as specified when creating the audience for uploading user IDs.
	 */
	IsIfaAudience bool `json:"isIfaAudience"`

	/**
	 * Get Permission
	 */
	Permission AudienceGroupPermission `json:"permission,omitempty"`

	/**
	 * Get CreateRoute
	 */
	CreateRoute AudienceGroupCreateRoute `json:"createRoute,omitempty"`
}