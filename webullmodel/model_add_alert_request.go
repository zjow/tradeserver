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
// AddAlertRequest struct for AddAlertRequest
type AddAlertRequest struct {
	DisExchangeCode string `json:"disExchangeCode,omitempty"`
	DisSymbol string `json:"disSymbol,omitempty"`
	EventWarningInput AddAlertRequestEventWarningInput `json:"eventWarningInput,omitempty"`
	ExchangeCode string `json:"exchangeCode,omitempty"`
	RegionId string `json:"regionId,omitempty"`
	ShowCode string `json:"showCode,omitempty"`
	TickerId int32 `json:"tickerId,omitempty"`
	TickerName string `json:"tickerName,omitempty"`
	TickerSymbol string `json:"tickerSymbol,omitempty"`
	TickerType string `json:"tickerType,omitempty"`
	TinyName string `json:"tinyName,omitempty"`
	WarningInput AddAlertRequestWarningInput `json:"warningInput,omitempty"`
}
