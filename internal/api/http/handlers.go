package http

import (
	"net/http"

	"github.com/CafeKetab/user/internal/models"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (s *Server) login(c *fiber.Ctx) error {
	request := new(models.UserCredential)
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		s.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	return nil
}

func (s *Server) register(c *fiber.Ctx) error {
	return nil
}
