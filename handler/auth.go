package handler

import (
	"gobasic/service"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const JwtSecret = "dsgsfdfnkjasnfjalsmdna"

type AuthHandler interface {
	SignUp(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
}

type authHandler struct {
	authSrv service.AuthService
}

func NewAuthHandler(authSrv service.AuthService) AuthHandler {
	return authHandler{authSrv: authSrv}
}

func (h authHandler) SignUp(c *fiber.Ctx) error {

	req := SignUpRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user, err := h.authSrv.SignUp(req.Email, req.Name, string(password))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "This email is already exist")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h authHandler) SignIn(c *fiber.Ctx) error {

	req := SignInRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	id, err := h.authSrv.SignIn(req.Email, req.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect email or password")
	}

	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(*id),
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(JwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"jwtToken": token,
	})
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
