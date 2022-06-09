package main

import (
	"gobasic/db"
	"gobasic/handler"
	"gobasic/repository"
	"gobasic/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	
	app := fiber.New()

	d := db.New()
	db.AutoMigrate(d)

	customerRepositoryDB := repository.NewCustomerRepositoryDB(d)
	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	app.Get("/customer", customerHandler.GetCustomers)
	app.Get("/customer/:id", customerHandler.GetCustomer)
	app.Post("/customer", customerHandler.CreateCustomer)
	app.Put("/customer/:id", customerHandler.UpdateCustomer)
	app.Delete("/customer/:id", customerHandler.DeleteCustomer)

	app.Listen(":8080")

}
