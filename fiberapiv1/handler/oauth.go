package handler

import (
	"encoding/json"
	"fiberapiv1/configs"
	"fiberapiv1/helper"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shareed2k/goth_fiber"
)

type OAuthHandler interface {
	LoginByProvider(c *fiber.Ctx) error
	ProviderCallBack(c *fiber.Ctx) error
	LineLogin(c *fiber.Ctx) error
	LineCallBack(c *fiber.Ctx) error
	GetLineToken(c *fiber.Ctx) error
	LogOut(c *fiber.Ctx) error
}

type oAuthHandler struct {
}

func NewOAuthHandler() oAuthHandler {
	return oAuthHandler{}
}

// GetUsers godoc
// @Description LoginByProvider with OAuth func
// @Tags OAuth2.0
// @Accept */*
// @Produce json
// @Param  provider path  string  true  "Provider example path /auth?provider=google""
// @response 200 "Success"
// @Router /v1/api/oauth/auth [get]
func (o oAuthHandler) LoginByProvider(c *fiber.Ctx) error {
	if gothUser, err := goth_fiber.CompleteUserAuth(c); err == nil {
		response := helper.BuildResponse(true, "LoginByProvider OAuth", gothUser)
		c.JSON(response)
	} else {
		goth_fiber.BeginAuthHandler(c)
	}
	return nil
}

// GetUsers godoc
// @Description ProviderCallBack callback func
// @Tags OAuth2.0
// @Accept */*
// @Produce json
// @response 200 "Success"
// @Router /v1/api/oauth/auth/callback  [get]
func (o oAuthHandler) ProviderCallBack(c *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		log.Fatal(err)
		return err
	}
	response := helper.BuildResponse(true, "Provider callback", user)
	return c.JSON(response)
}

// GetUsers godoc
// @Description LineLogin with OAuth func
// @Tags OAuth2.0
// @Accept */*
// @Produce json
// @response 200 "Success"
// @Router /v1/api/oauth/auth/line [get]
func (o oAuthHandler) LineLogin(c *fiber.Ctx) error {
	callbackURL := "http://localhost:8000/v1/api/oauth/auth/line/callback"
	uri := fmt.Sprintf("https://access.line.me/oauth2/v2.1/authorize?response_type=code&state=xxx&client_id=%s&redirect_uri=%s&scope=profile openid", configs.GetLineClientID(), callbackURL)
	return c.Redirect(uri)
}

// GetUsers godoc
// @Description LineCallBack callback func
// @Tags OAuth2.0
// @Accept */*
// @Produce json
// @response 200 "Success"
// @Router /v1/api/oauth/auth/line/callback  [get]
func (o oAuthHandler) LineCallBack(c *fiber.Ctx) error {
	queryValue := c.Query("code")
	data := map[string]interface{}{
		"code": queryValue,
	}
	response := helper.BuildResponse(true, "oAuth callback", data)
	return c.JSON(response)
}

// GetUsers godoc
// @Description GetLineToken callback func
// @Tags OAuth2.0
// @Accept */*
// @Produce json
// @response 200 "Success"
// @Router /v1/api/oauth/auth/line/token  [get]
func (o oAuthHandler) GetLineToken(c *fiber.Ctx) error {
	var uri = "https://api.line.me/oauth2/v2.1/token"
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("code", c.FormValue("code"))
	data.Set("redirect_uri", "http://localhost:8000/v1/api/oauth/auth/line/callback")
	data.Set("client_id", configs.GetLineClientID())
	data.Set("client_secret", configs.GetLineClientSecret())
	encodedData := data.Encode()
	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(encodedData))
	//req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBufferString(encodedData))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	/// http client end ///
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	return c.JSON(res)
}

// GetUsers godoc
// @Description OAuth logout func
// @Tags OAuth2.0
// @Accept */*
// @Produce json
// @response 200 "Success"
// @Router /v1/api/oauth/logout  [get]
func (o oAuthHandler) LogOut(c *fiber.Ctx) error {
	if err := goth_fiber.Logout(c); err != nil {
		log.Fatal(err)
	}
	return c.SendString("logout")
}
