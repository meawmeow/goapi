package routes

import (
	"fiberapiv1/configs"
	"fiberapiv1/handler"
	"fiberapiv1/repository"
	"fiberapiv1/security"
	"fiberapiv1/services"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutesSetup(app *fiber.App) {

	db := configs.Database.Db
	userRepo := repository.NewUserRepositoryDB(db)
	// userRepoMock repository.UserRepository = repository.NewUserRepositoryMock(db)
	userSrv := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSrv)

	v1 := app.Group("/v1/api")
	//public
	v1.Post("/register", userHandler.Register)
	v1.Post("/login", userHandler.Login)

	//private use auth
	//v1.Use(security.AuthorizationRequired())
	v1.Get("/users", security.AuthorizationRequired(), userHandler.GetUsers)
	v1.Get("/user", security.AuthorizationRequired(), userHandler.GetUser)
	v1.Post("/logout", security.AuthorizationRequired(), userHandler.LogOut)
	v1.Post("/uploadfile", security.AuthorizationRequired(), userHandler.UploadFile)
	v1.Get("/imageprofile", security.AuthorizationRequired(), userHandler.GetImageProfile)
}
