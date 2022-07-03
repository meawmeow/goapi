package main

import (
	"fiberapiv1/configs"
	"fiberapiv1/routes"
	"fiberapiv1/security"

	_ "fiberapiv1/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func init() {
	configs.InitTimeZone()
	configs.InitConfig()
	configs.InitDatabase()
}

// @title Fiber KiwShop API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:8000
// @BasePath /
func main() {

	app := fiber.New(fiber.Config{
		BodyLimit: 4 * 1024 * 1024, // this is the default limit of 4MB
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Use(recover.New())
	app.Use(cors.New())

	routes.AuthRoutesSetup(app)
	routes.OAuthProviderRoutesSetup(app)
	//Private
	app.Use(security.AuthorizationRequired())
	routes.ProductRoutesSetup(app)
	routes.OrdersRoutesSetup(app)

	app.Listen((":8000"))
}
