package handler

import (
	"fyndr.com/api/src/internal/service"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) Handler {
	return &AuthHandler{
		authService: authService,
	}
}

func (a AuthHandler) GetAuth(c *fiber.Ctx) error {
	redirectUrl, _ := a.authService.Login("test")
	return c.Redirect(redirectUrl)
}

func (a AuthHandler) GetCallback(c *fiber.Ctx) error {
	return c.SendString("Callback")
}

func (a AuthHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/login", a.GetAuth)
	router.Get("/callback", a.GetCallback)
}
