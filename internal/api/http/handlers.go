package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (s *Server) register(c *fiber.Ctx) error {
	ctx := c.Context()

	request := struct{ Email, Password string }{}
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		s.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	user, err := s.repository.FindUserByEmail(ctx, request.Email)
	if err != nil {
		return err
	}

	if user.Id != 0 {
		// user already exists
		return err
	}

	user.Password = request.Password
	if err := s.repository.CreateUser(ctx, user); err != nil {
		return err
	}

	// request token

	return nil
}

func (s *Server) login(c *fiber.Ctx) error {
	ctx := c.Context()

	request := struct{ Email, Password string }{}
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		s.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	user, err := s.repository.FindUserByEmailAndPassword(ctx, request.Email, request.Password)
	if err != nil {
		return err
	}

	// check nil
	if user == nil {
		// invalid email or password (doesn't find any)
		return nil
	}

	// request token

	return nil
}

func (s *Server) update(c *fiber.Ctx) error {
	return nil
}
