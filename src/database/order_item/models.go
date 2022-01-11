// Code generated by sqlc. DO NOT EDIT.

package order_item

import ()

type OrderItem struct {
	ID                int64   `json:"id"`
	OrderID           int64   `json:"order_id"`
	ProductTitle      string  `json:"product_title"`
	Price             float64 `json:"price"`
	Quantity          int64   `json:"quantity"`
	AdminRevenue      float64 `json:"admin_revenue"`
	AmbassadorRevenue float64 `json:"ambassador_revenue"`
}