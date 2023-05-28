package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"inventory-api/src/database"
	"inventory-api/src/middleware"
	"inventory-api/src/routes/auth"
	"inventory-api/src/routes/clones"
	"inventory-api/src/routes/customers"
	"inventory-api/src/routes/orders"
	"inventory-api/src/routes/users"
	"os"
)

var PORT = ":3000"

func main() {
	possiblePortEnv := os.Getenv("PORT")
	if possiblePortEnv != "" {
		PORT = ":" + os.Getenv("PORT")
	}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(helmet.New())

	database.Connect()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("mongo", database.Client.Database("inventory"))
		return c.Next()
	})

	authRoute := app.Group("/auth")
	{
		authRoute.Post("/login", auth.Login)
	}

	usersRoute := app.Group("/users")
	{
		usersRoute.Get("/:apikey", users.GetByAPIKey)
		usersRoute.Get("/", users.GetAll).Use(middleware.APIKeyAuth)
		usersRoute.Post("/create", users.Create).Use(middleware.APIKeyAuth)
		usersRoute.Delete("/:id/delete", users.Delete).Use(middleware.APIKeyAuth)
		usersRoute.Put("/:id/admin", users.ToggleAdmin).Use(middleware.APIKeyAuth)
	}

	clonesRoute := app.Group("/clones").Use(middleware.APIKeyAuth)
	{
		clonesRoute.Get("/", clones.GetAll)
		clonesRoute.Post("/create", clones.Create)
		clonesRoute.Delete("/:id/delete", clones.Delete)
		clonesRoute.Put("/:id/quantity/:quantity", clones.SetQuantity)
		clonesRoute.Put("/:id/name", clones.EditName)
	}

	ordersRoute := app.Group("/orders").Use(middleware.APIKeyAuth)
	{
		ordersRoute.Get("/", orders.GetAll)
		ordersRoute.Post("/place", orders.Place)
		ordersRoute.Put("/:id/status/", orders.SetStatus)
		ordersRoute.Put("/:id/notes", orders.UpdateNotes)
	}

	customersRoute := app.Group("/customers").Use(middleware.APIKeyAuth)
	{
		customersRoute.Get("/", customers.GetAll)
		customersRoute.Post("/create", customers.Create)
		customersRoute.Delete("/:id/delete", customers.Delete)
		customersRoute.Put("/:id/notes", customers.UpdateNotes)
		customersRoute.Get("/names", customers.Names)
	}

	err := app.Listen("0.0.0.0" + PORT)
	if err != nil {
		panic(err)
	}
}
