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

// GetAudienceGroupsResponse
// Gets data for more than one audience.
// https://developers.line.biz/en/reference/messaging-api/#get-audience-groups
type GetAudienceGroupsResponse struct {

	/**
	 * An array of audience data. If there are no audiences that match the specified filter, an empty array will be returned.
	 */
	AudienceGroups []AudienceGroup `json:"audienceGroups,omitempty"`

	/**
	 * true when this is not the last page.
	 */
	HasNextPage bool `json:"hasNextPage"`

	/**
	 * The total number of audiences that can be returned with the specified filter.
	 */
	TotalCount int64 `json:"totalCount"`

	/**
	 * Of the audiences you can get with the specified filter, the number of audiences with the update permission set to READ_WRITE.
	 */
	ReadWriteAudienceGroupTotalCount int64 `json:"readWriteAudienceGroupTotalCount"`

	/**
	 * The current page number.
	 */
	Page int64 `json:"page"`

	/**
	 * The maximum number of audiences on the current page.
	 */
	Size int64 `json:"size"`
}
