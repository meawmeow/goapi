package helper

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var CookieTokenName = "jwt"

func CreateTokenCookie(token string, c *fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:     CookieTokenName,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
}
func ClearTokenCookie(c *fiber.Ctx) {
	cookie := fiber.Cookie{
		Name:     CookieTokenName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
}
