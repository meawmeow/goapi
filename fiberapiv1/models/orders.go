package models

import "fiberapiv1/entity"

type OrdersRequest struct {
	Count     int `json:"count"`
	ProductID int `json:"product_id"`
	UserID    int `json:"user_id"`
}
type OrdersResponse struct {
	ID        int            `json:"id"`
	Count     int            `json:"count"`
	NetPrices uint64         `json:"net_price"`
	Product   entity.Product `json:"Product"`
	User      entity.User    `json:"User"`
}
