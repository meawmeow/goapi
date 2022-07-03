package routes

import (
	"fiberapiv1/configs"
	"fiberapiv1/handler"
	"fiberapiv1/repository"
	"fiberapiv1/services"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutesSetup(app *fiber.App) {

	db := configs.Database.Db
	proRepo := repository.NewProductRepositoryDB(db)
	proSrv := services.NewProductService(proRepo)
	proHandler := handler.NewProducrHandler(proSrv)

	v1 := app.Group("/v1/api")
	//private use auth
	//v1.Use(security.AuthorizationRequired())
	v1.Post("/productgroup", proHandler.CreateProductGroup)
	v1.Post("/product", proHandler.CreateProduct)
	v1.Get("/productgroups", proHandler.ProductGroups)
	v1.Get("/products", proHandler.Products)
	v1.Get("/products/group", proHandler.ProductsByGroupId)
}
