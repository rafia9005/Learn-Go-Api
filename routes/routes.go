package routes

import (
	"Learn-Go-Api/handler"
	"Learn-Go-Api/middleware"
	"Learn-Go-Api/model/entity"

	"github.com/gofiber/fiber/v2"
)

//middleware Auth & Admin Role
var Auth = middleware.Auth
var Admin = middleware.AdminRole

func SetupRouter(app *fiber.App){
  //routing
  app.Static("/", "./public")
  app.Get("/example", handler.GetExample)
  app.Get("/example/:id", handler.GetByIdExample)
  app.Post("/example", handler.CreateExample)
  app.Delete("/example/:id", handler.DeleteExample)
  app.Put("/example/:id", handler.UpdateExample)
}

func AutoMigrate(){
  //run migration
  RunMigrate(&entity.Users{})
  RunMigrate(&entity.Example{})
}
