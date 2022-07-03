package routes

import (
	"fiberapiv1/configs"
	"fiberapiv1/handler"
	"fiberapiv1/repository"
	"fiberapiv1/services"

	"github.com/gofiber/fiber/v2"
)

func OrdersRoutesSetup(app *fiber.App) {

	db := configs.Database.Db
	ordersRepo := repository.NewOrdersRepositoryDB(db)
	ordersSrv := services.NewOrdersService(ordersRepo)
	ordersHandler := handler.NewOrdersHandler(ordersSrv)

	v1 := app.Group("/v1/api")
	//private use auth
	//v1.Use(security.AuthorizationRequired())
	v1.Post("/orders", ordersHandler.CreateOrders)
	v1.Get("/ordersall", ordersHandler.GetAllOrders)
	v1.Get("/ordersbyuserid", ordersHandler.GetOrdersByUserId)
	v1.Delete("/orders", ordersHandler.DeleteOrdersById)
}
