package middleware

import (
	"errors"
	"fmt"
	"bitbucket.org/friendsonly/core/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

// Protected protect routes
func Protected(c *fiber.Ctx) error {
	if c.Locals("user_id") != nil {
			return c.Next()
	}

	return jwtError(c, errors.New("Missing or malformed JWT "))
}

// Protected protect routes
func Authenticate() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Config("JWT_SECRET")),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Next()
		},
		SuccessHandler: jwtSuccess,
	})
}

func jwtSuccess(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	for k, v := range claims {
		c.Locals(
			k,
			fmt.Sprintf("%v",v),
		)
	}
	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

