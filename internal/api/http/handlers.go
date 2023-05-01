package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (handler *Server) register(c *fiber.Ctx) error {
	ctx := c.Context()

	request := struct{ Email, Password string }{}
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	user, err := handler.repository.FindUserByEmail(ctx, request.Email)
	if err != nil {
		return err
	}

	if user.Id != 0 {
		// user already exists
		return err
	}

	user.Password = request.Password
	if err := handler.repository.CreateUser(ctx, user); err != nil {
		return err
	}

	// request token
	token, err := handler.auth.GenerateToken(ctx, user.Id)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).SendString(token)
}

func (handler *Server) login(c *fiber.Ctx) error {
	ctx := c.Context()

	request := struct{ Email, Password string }{}
	if err := c.BodyParser(&request); err != nil {
		errString := "Error parsing request body"
		handler.logger.Error(errString, zap.Error(err))
		return c.Status(http.StatusBadRequest).SendString(errString)
	}

	user, err := handler.repository.FindUserByEmailAndPassword(ctx, request.Email, request.Password)
	if err != nil {
		return err
	}

	// check nil
	if user == nil {
		// invalid email or password (doesn't find any)
		return nil
	}

	// request token
	token, err := handler.auth.GenerateToken(ctx, user.Id)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).SendString(token)
}

// get user by id
func (handler *Server) user(c *fiber.Ctx) error {
	c.Params("id")
	return nil
}

// get user of the header
func (handler *Server) me(c *fiber.Ctx) error {
	return nil
}

func (handler *Server) update(c *fiber.Ctx) error {
	return nil
}
