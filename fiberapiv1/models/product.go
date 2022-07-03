package models

import "fiberapiv1/entity"

type ProductRequest struct {
	Name           string `json:"name"`
	Price          string `json:"price"`
	ImageUri       string `json:"imageuri"`
	ProductGroupID int    `json:"productGroup"`
}
type ProductResponse struct {
	ID           int                 `json:"id"`
	Name         string              `json:"name"`
	Price        string              `json:"price"`
	ImageUri     string              `json:"imageuri"`
	ProductGroup entity.ProductGroup `json:"productGroup"`
}
