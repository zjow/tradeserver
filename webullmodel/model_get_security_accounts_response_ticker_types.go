/*
 * Webull API
 *
 * Webull API Documentation
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package webullmodel
// GetSecurityAccountsResponseTickerTypes struct for GetSecurityAccountsResponseTickerTypes
type GetSecurityAccountsResponseTickerTypes struct {
	OrderTypes []OrderType `json:"orderTypes,omitempty"`
	RegionId int32 `json:"regionId,omitempty"`
	TickerType int32 `json:"tickerType,omitempty"`
}
