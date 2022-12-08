package routes

import (
	controllers "go-fiber-test/controllers"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hellowwwww")
	// })

	app.Post("/", controllers.HelloWorld)

	app.Get("/", controllers.HelloWorld)

	app.Post("/", controllers.Name)

	app.Get("/user/:name", controllers.GetUser)

	app.Post("/inet", controllers.Search)

	// Parameters
	app.Get("/user/:name/books/:title", controllers.Params)

	api := app.Group("/api")               // /api
	v1 := api.Group("/v1")                 // /api/v1
	v1.Get("/list", controllers.GetList)   // /api/v1/list
	v1.Get("/user", controllers.V1GetUser) // /api/v1/user
	v1.Get("/fact/:n", controllers.Fact)   // /api/v1/fact/:n

	v2 := api.Group("/v2")              // /api/v2
	v2.Get("/list", controllers.V2List) // /api/v2/list
	v2.Get("/user", controllers.V2User) // /api/v2/user

	v2.Post("/adduser", controllers.AddUser)

	//CRUD
	v1.Get("/dog", controllers.GetDogs)
	v1.Get("/dog/filter", controllers.GetDog)
	v1.Post("/dog", controllers.AddDog)
	v1.Put("/dog/:id", controllers.UpdateDog)
	v1.Delete("/dog/:id", controllers.RemoveDog)

}
