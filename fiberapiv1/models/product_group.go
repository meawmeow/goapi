package models

type ProductGroupRequest struct {
	Name     string `json:"name"`
	ImageUri string `json:"imageuri"`
}
type ProductGroupResponse struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	ImageUri string `json:"imageuri"`
}
