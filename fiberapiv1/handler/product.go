package handler

import (
	"fiberapiv1/errs"
	"fiberapiv1/helper"
	"fiberapiv1/models"
	"fiberapiv1/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler interface {
	CreateProduct(c *fiber.Ctx) error
	CreateProductGroup(c *fiber.Ctx) error
	Products(c *fiber.Ctx) error
	ProductsByGroupId(c *fiber.Ctx) error
	ProductGroups(c *fiber.Ctx) error
}

type productHandler struct {
	productSrv services.ProductService
}

func NewProducrHandler(productSrv services.ProductService) ProductHandler {
	return productHandler{productSrv: productSrv}
}

// CreateProduct godoc
// @Description Create Product func
// @Tags Product
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @param Product body models.ProductRequest true "data to be insert product"
// @response 200 {object} helper.Response "OK"
// @Router /v1/api/product  [post]
func (h productHandler) CreateProduct(c *fiber.Ctx) error {
	//filter param for handler
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		errRes := helper.BuildErrorResponse("path CreateProduct", err, nil)
		return c.JSON(errRes)
	}
	groupId, err := strconv.Atoi(data["groupID"])
	if err != nil {
		return err
	}
	request := models.ProductRequest{
		Name:           data["name"],
		Price:          data["price"],
		ImageUri:       data["imageuri"],
		ProductGroupID: groupId,
	}
	//business logic for service
	err = h.productSrv.CreateProduct(request)
	if err != nil {
		errRes := helper.BuildErrorResponse("path CreateProduct", err, nil)
		return c.JSON(errRes)
	}
	//respone for register
	response := helper.BuildResponse(true, "CreateProduct success", nil)
	return c.JSON(response)
}

// CreateProductGroup godoc
// @Description Create ProductGroup func
// @Tags Product
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @param ProductGroup body models.ProductGroupRequest true "data to be insert product group"
// @response 200 {object} helper.Response "OK"
// @Router /v1/api/productgroup  [post]
func (h productHandler) CreateProductGroup(c *fiber.Ctx) error {
	//filter param for handler
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		errRes := helper.BuildErrorResponse("path CreateProductGroup", err, nil)
		return c.JSON(errRes)
	}
	request := models.ProductGroupRequest{
		Name:     data["name"],
		ImageUri: data["imageuri"],
	}
	//business logic for service
	err := h.productSrv.CreateProductGroup(request)
	if err != nil {
		errRes := helper.BuildErrorResponse("path CreateProductGroup", err, nil)
		return c.JSON(errRes)
	}
	//respone for register
	response := helper.BuildResponse(true, "CreateProductGroup success", nil)
	return c.JSON(response)
}

// ProductGroups godoc
// @Description Get All ProductGroups func
// @Tags Product
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @response 200 {object} helper.Response "OK"
// @Router /v1/api/productgroups  [get]
func (h productHandler) ProductGroups(c *fiber.Ctx) error {

	pGroups, err := h.productSrv.ProductGroups()
	if err != nil {
		errRes := helper.BuildErrorResponse("path ProductGroups", err, nil)
		return c.JSON(errRes)
	}
	response := helper.BuildResponse(true, "get all ProductGroups", pGroups)
	return c.JSON(response)
}

// Products godoc
// @Description Get All Products func
// @Tags Product
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @response 200 {object} helper.Response "OK"
// @Router /v1/api/products  [get]
func (h productHandler) Products(c *fiber.Ctx) error {

	products, err := h.productSrv.Products()
	if err != nil {
		errRes := helper.BuildErrorResponse("path Products", err, nil)
		return c.JSON(errRes)
	}
	response := helper.BuildResponse(true, "get all Products", products)
	return c.JSON(response)
}

// ProductsByGroupId godoc
// @Description Get Products By GroupId func
// @Tags Product
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Param groupId path int64 true "groupId"
// @response 200 "Success"
// @Router /v1/api/products/group  [get]
func (h productHandler) ProductsByGroupId(c *fiber.Ctx) error {

	if s := c.Query("groupId"); s == "" {
		errRes := helper.BuildErrorResponse("path ProductsByGroupId", errs.NewError("parsing param invalid"), nil)
		return c.JSON(errRes)
	}
	groupId, err := strconv.Atoi(c.Query("groupId"))
	if err != nil {
		errRes := helper.BuildErrorResponse("path ProductsByGroupId", errs.NewError("parsing param invalid syntax"), nil)
		return c.JSON(errRes)
	}

	products, err := h.productSrv.ProductsByGroupId(groupId)
	if err != nil {
		errRes := helper.BuildErrorResponse("path Products", err, nil)
		return c.JSON(errRes)
	}
	response := helper.BuildResponse(true, "get all Products", products)
	return c.JSON(response)
}
