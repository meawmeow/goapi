package handler

import (
	"fiberapiv1/errs"
	"fiberapiv1/helper"
	"fiberapiv1/models"
	"fiberapiv1/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	LogOut(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	UploadFile(c *fiber.Ctx) error
	GetImageProfile(c *fiber.Ctx) error
}

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) UserHandler {
	return userHandler{userSrv: userSrv}
}

// Register godoc
// @Description Register func
// @Tags Authorization
// @Accept */*
// @Produce json
// @param User body models.UserRequest true "user data to be register"
// @response 200 {object} helper.Response "OK"
// @Router /v1/api/register  [post]
func (h userHandler) Register(c *fiber.Ctx) error {
	//filter param for handler
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		errRes := helper.BuildErrorResponse("path register", err, nil)
		return c.JSON(errRes)
	}
	request := models.UserRequest{
		Username: data["username"],
		Email:    data["email"],
		Password: data["password"],
		Address:  data["address"],
	}
	//business logic for service
	user, err := h.userSrv.CreateUser(request)
	if err != nil {
		errRes := helper.BuildErrorResponse("path register", err, nil)
		return c.JSON(errRes)
	}
	//respone for user
	response := helper.BuildResponse(true, "register success", user)
	return c.JSON(response)
}

// Login godoc
// @Description Login func
// @Tags Authorization
// @Accept json
// @Produce json
// @param User body models.UserRequest true "user data to be login"
// @response 200 {object} models.UserResponse "OK"
// @Router /v1/api/login  [post]
func (h userHandler) Login(c *fiber.Ctx) error {
	//filter param for handler
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		errRes := helper.BuildErrorResponse("path login", err, nil)
		return c.JSON(errRes)
	}
	request := models.UserRequest{
		Username: data["username"],
		Email:    data["email"],
		Password: data["password"],
	}
	//business logic for service
	user, err := h.userSrv.Login(request)
	if err != nil {
		errRes := helper.BuildErrorResponse("path login", err, nil)
		return c.JSON(errRes)
	}
	helper.CreateTokenCookie(user.Token, c)
	//respone for user
	response := helper.BuildResponse(true, "login success", user)
	return c.JSON(response)
}

// LogOut godoc
// @Description Logout func
// @Tags Authorization
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @response 200 "Success"
// @Router /v1/api/logout  [post]
func (h userHandler) LogOut(c *fiber.Ctx) error {
	helper.ClearTokenCookie(c)
	response := helper.BuildResponse(true, "logout success", nil)
	return c.JSON(response)
}

// GetUsers godoc
// @Description GetUser func
// @Tags Authorization
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @response 200 "Success"
// @Router /v1/api/users  [get]
func (h userHandler) GetUsers(c *fiber.Ctx) error {

	users, err := h.userSrv.GetUsers()
	if err != nil {
		errRes := helper.BuildErrorResponse("path users", err, nil)
		return c.JSON(errRes)
	}
	response := helper.BuildResponse(true, "get all users", users)
	return c.JSON(response)
}

// GetUser godoc
// @Description GetUser func
// @Tags Authorization
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @Param id query int64 true "userId"
// @response 200 "Success"
// @Router /v1/api/user/{id}  [get]
func (h userHandler) GetUser(c *fiber.Ctx) error {
	if s := c.Query("id"); s == "" {
		errRes := helper.BuildErrorResponse("path user", errs.NewError("parsing param invalid"), nil)
		return c.JSON(errRes)
	}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errRes := helper.BuildErrorResponse("path user", errs.NewError("parsing param invalid syntax"), nil)
		return c.JSON(errRes)
	}

	userRespones, err := h.userSrv.GetUser(id)
	if err != nil {
		errRes := helper.BuildErrorResponse("path user", err, nil)
		return c.JSON(errRes)
	}
	response := helper.BuildResponse(true, "get user by Id : "+c.Query("id"), userRespones)
	return c.JSON(response)
}

// GetUsers godoc
// @Description MultipartForm files
// @Tags Authorization
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @response 200 "Success"
// @Router /v1/api/uploadfile  [post]
func (h userHandler) UploadFile(c *fiber.Ctx) error {
	name := c.FormValue("name")
	file, err := c.FormFile("fileUpload")
	if err != nil {
		errRes := helper.BuildErrorResponse("Error UploadFile imag type", errs.NewError(err.Error()), nil)
		return c.JSON(errRes)
	}
	if err := c.SaveFile(file, fmt.Sprintf("assets/upload/%s", file.Filename)); err != nil {
		errRes := helper.BuildErrorResponse("Error save file failed", errs.NewError(err.Error()), nil)
		return c.JSON(errRes)
	}
	response := helper.BuildResponse(true, "UploadFile Successed", name)
	return c.JSON(response)
}

// GetUsers godoc
// @Description GetImageProfile func
// @Tags Authorization
// @Accept */*
// @Produce json
// @Security ApiKeyAuth
// @response 200 "Success"
// @Router /v1/api/imageprofile  [get]
func (h userHandler) GetImageProfile(c *fiber.Ctx) error {

	// response := helper.BuildResponse(true, "image profile Successed", nil)
	// return c.JSON(response)
	//return c.SendFile("assets/images/pokemon1.png")
	return c.Download("assets/images/pokemon1.png")
}
