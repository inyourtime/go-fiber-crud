package handler

import (
	"gobasic/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler interface {
	GetCustomers(c *fiber.Ctx) error
	GetCustomer(c *fiber.Ctx) error
	CreateCustomer(c *fiber.Ctx) error
	UpdateCustomer(c *fiber.Ctx) error
	DeleteCustomer(c *fiber.Ctx) error
}

type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) CustomerHandler {
	return customerHandler{custSrv: custSrv}
}

func (h customerHandler) GetCustomers(c *fiber.Ctx) error {

	customers, err := h.custSrv.GetCustomers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"status":    "ok",
		"customers": customers,
	})
}

func (h customerHandler) GetCustomer(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	customer, err := h.custSrv.GetCustomer(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"customer": &customer,
	})
}

func (h customerHandler) CreateCustomer(c *fiber.Ctx) error {

	custRequest := CustomerRequest{}

	err := c.BodyParser(&custRequest)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(&custRequest)
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}
	// fmt.Println(custRequest)

	customer, err := h.custSrv.CreateCustomer(custRequest.Name, custRequest.City, custRequest.ZipCode)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":   "ok",
		"customer": customer,
	})
}

func (h customerHandler) UpdateCustomer(c *fiber.Ctx) error {

	custRequest := CustomerUpdateRequest{}

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	err = c.BodyParser(&custRequest)
	if err != nil {
		return err
	}

	err = h.custSrv.UpdateCustomer(id, custRequest.Name, custRequest.City, custRequest.ZipCode)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "updated",
	})
}

func (h customerHandler) DeleteCustomer(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrUnprocessableEntity
	}

	err = h.custSrv.DeleteCustomer(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "deleted",
	})
}

type CustomerRequest struct {
	Name    string `json:"name" validate:"required"`
	City    string `json:"city" validate:"required"`
	ZipCode string `json:"zipcode" validate:"required"`
}

type CustomerUpdateRequest struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
}
