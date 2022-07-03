package handler

import (
	"fiberapiv1/errs"
	"fiberapiv1/helper"
	"fiberapiv1/models"
	"fiberapiv1/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OrdersHandler interface {
	CreateOrders(c *fiber.Ctx) error
	GetAllOrders(c *fiber.Ctx) error
	GetOrdersByUserId(c *fiber.Ctx) error
	DeleteOrdersById(c *fiber.Ctx) error
}

type ordersHandler struct {
	ordersSrv services.OrdersService
}

func NewOrdersHandler(ordersSrv services.OrdersService) OrdersHandler {
	return ordersHandler{ordersSrv: ordersSrv}
}

// CreateOrders godoc
// @Description Create Orders func
// @Tags Orders
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @param Orders body models.OrdersRequest true "data to be insert orders"
// @response 200 {object} helper.Response "OK"
// @Router /v1/api/orders  [post]
func (h ordersHandler) CreateOrders(c *fiber.Ctx) error {
	//filter param for handler
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		errRes := helper.BuildErrorResponse("path CreateProduct", err, nil)
		return c.JSON(errRes)
	}
	count, err := strconv.Atoi(data["count"])
	if err != nil {
		return err
	}
	product_id, err := strconv.Atoi(data["product_id"])
	if err != nil {
		return err
	}
	user_id, err := strconv.Atoi(data["user_id"])
	if err != nil {
		return err
	}
	request := models.OrdersRequest{
		Count:     count,
		ProductID: product_id,
		UserID:    user_id,
	}
	//business logic for service
	err = h.ordersSrv.CreateOrders(request)
	if err != nil {
		errRes := helper.BuildErrorResponse("path CreateProduct", err, nil)
		return c.JSON(errRes)
	}
	//respone for register
	response := helper.BuildResponse(true, "CreateProduct success", nil)
	return c.JSON(response)
}

// GetAllOrders godoc
// @Description Get All Orders func
// @Tags Orders
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @response 200 {object} helper.Response "OK"
// @Router /v1/api/ordersall  [get]
func (h ordersHandler) GetAllOrders(c *fiber.Ctx) error {

	orders, err := h.ordersSrv.OrdersAll()
	if err != nil {
		errRes := helper.BuildErrorResponse("path GetAllOrders", err, nil)
		return c.JSON(errRes)
	}
	//respone for register
	response := helper.BuildResponse(true, "GetAllOrders success", orders)
	return c.JSON(response)
}

// GetOrdersByUserId godoc
// @Description Get Orders By UserId func
// @Tags Orders
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Param userId query int64 true "userId"
// @response 200 {object} helper.Response "OK"
// @Router /v1/api/ordersbyuserid  [get]
func (h ordersHandler) GetOrdersByUserId(c *fiber.Ctx) error {

	if s := c.Query("userId"); s == "" {
		errRes := helper.BuildErrorResponse("path GetOrdersByUserId", errs.NewError("parsing param invalid"), nil)
		return c.JSON(errRes)
	}
	userId, err := strconv.Atoi(c.Query("userId"))
	if err != nil {
		errRes := helper.BuildErrorResponse("path GetOrdersByUserId", errs.NewError("parsing param invalid syntax"), nil)
		return c.JSON(errRes)
	}

	orders, err := h.ordersSrv.OrdersByUserId(userId)
	if err != nil {
		errRes := helper.BuildErrorResponse("path GetOrdersByUserId", err, nil)
		return c.JSON(errRes)
	}

	response := helper.BuildResponse(true, "GetOrdersByUserId success", orders)
	return c.JSON(response)
}

// GetOrdersByUserId godoc
// @Description Get Orders By UserId func
// @Tags Orders
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Param id query int64 true "id"
// @response 200 {object} helper.Response "OK"
// @Router /v1/api/orders  [delete]
func (h ordersHandler) DeleteOrdersById(c *fiber.Ctx) error {

	if s := c.Query("id"); s == "" {
		errRes := helper.BuildErrorResponse("path DeleteOrdersById", errs.NewError("parsing param invalid"), nil)
		return c.JSON(errRes)
	}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := helper.BuildErrorResponse("path DeleteOrdersById", errs.NewError("parsing param invalid syntax"), nil)
		return c.JSON(errRes)
	}

	err = h.ordersSrv.DeleteOrdersById(id)
	if err != nil {
		errRes := helper.BuildErrorResponse("path DeleteOrdersById", err, nil)
		return c.JSON(errRes)
	}
	response := helper.BuildResponse(true, "DeleteOrdersById success", nil)
	return c.JSON(response)
}
