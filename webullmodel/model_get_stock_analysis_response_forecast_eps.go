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
// GetStockAnalysisResponseForecastEps struct for GetStockAnalysisResponseForecastEps
type GetStockAnalysisResponseForecastEps struct {
	CurrencyName string `json:"currencyName,omitempty"`
	Id string `json:"id,omitempty"`
	Points []GetStockAnalysisResponseForecastEpsPoints `json:"points,omitempty"`
	Title string `json:"title,omitempty"`
}
