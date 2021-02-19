/*
 * Webull API
 *
 * Webull API Documentation
 *
 * API version: 3.0.1
 * Contact: austin.millan@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// GetOrdersItem struct for GetOrdersItem
type GetOrdersItem struct {
	Action OrderSide `json:"action,omitempty"`
	ComboId string `json:"comboId,omitempty"`
	ComboTickerType string `json:"comboTickerType,omitempty"`
	ComboType string `json:"comboType,omitempty"`
	FilledQuantity string `json:"filledQuantity,omitempty"`
	LmtPrice string `json:"lmtPrice,omitempty"`
	Orders []GetOrdersItemOrders `json:"orders,omitempty"`
	OutsideRegularTradingHour bool `json:"outsideRegularTradingHour,omitempty"`
	Quantity string `json:"quantity,omitempty"`
	Status OrderStatus `json:"status,omitempty"`
	TimeInForce string `json:"timeInForce,omitempty"`
}