package routes

import (
	"fiberapiv1/configs"
	"fiberapiv1/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func OAuthProviderRoutesSetup(app *fiber.App) {

	goth.UseProviders(
		google.New(
			configs.GetGoogleClientID(),
			configs.GetGoogleClientSecret(),
			"http://localhost:8000/v1/api/oauth/auth/callback"), //or http://localhost:8000/v1/api/oauth/auth/google/callback
		// line.New(
		// 	"1657138695",
		// 	"05528203f76fc3e126316cbecd84cb49",
		// 	"http://localhost:8000/v1/api/oauth/auth/line/callback", "profile", "openid", "email"),
	)
	oAuthHandler := handler.NewOAuthHandler()
	v1 := app.Group("/v1/api/oauth")
	//v1.Get("/auth", goth_fiber.BeginAuthHandler)
	//with lib goth multi provider
	v1.Get("/auth", oAuthHandler.LoginByProvider)
	v1.Get("/auth/callback", oAuthHandler.ProviderCallBack)
	//with line uri provider
	v1.Get("/auth/line", oAuthHandler.LineLogin)
	v1.Get("/auth/line/callback", oAuthHandler.LineCallBack)
	v1.Post("/auth/line/token", oAuthHandler.GetLineToken)

	v1.Get("/logout", oAuthHandler.LogOut)

}
