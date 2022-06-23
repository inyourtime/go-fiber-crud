package main

import (
	"gobasic/database"
	"gobasic/handler"
	"gobasic/repository"
	"gobasic/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)


func main() {

	app := fiber.New()

	app.Use(logger.New())
	app.Use("/hello", jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(handler.JwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	}))

	db := database.New()
	database.AutoMigrate(db)

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	app.Get("/customer", customerHandler.GetCustomers)
	app.Get("/customer/:id", customerHandler.GetCustomer)
	app.Post("/customer", customerHandler.CreateCustomer)
	app.Put("/customer/:id", customerHandler.UpdateCustomer)
	app.Delete("/customer/:id", customerHandler.DeleteCustomer)

	userRepositoryDB := repository.NewUserRepositoryDB(db)
	authService := service.NewAuthService(userRepositoryDB)
	authHandler := handler.NewAuthHandler(authService)

	app.Post("/sign-up", authHandler.SignUp)
	app.Post("/sign-in", authHandler.SignIn)

	app.Listen(":5000")

}
