package middleware

import (
	utils "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
)

func ValidateToken(c *fiber.Ctx) error {
	token := c.Get(utils.AUTHORIZATION)

	if len(token) > 7 && token[:7] == utils.BEARER {
		return c.Next()

	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		utils.STATUS:  fiber.StatusUnauthorized,
		utils.MESSAGE: utils.TOKEN_INVALID,
	})
}
