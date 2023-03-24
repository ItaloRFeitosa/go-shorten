package fiberutil

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-shorten/pkg/errs"
	"github.com/italorfeitosa/go-shorten/pkg/validate"
)

func Body(c *fiber.Ctx, v any) error {
	if err := c.BodyParser(v); err != nil {
		return err
	}

	return validate.Struct(v)
}

func Query(c *fiber.Ctx, v any) error {
	if err := c.QueryParser(v); err != nil {
		return err
	}

	return validate.Struct(v)
}

func Data(data any) fiber.Map {
	return fiber.Map{
		"data": data,
	}
}

func Error(err error) fiber.Map {
	return fiber.Map{
		"error": err,
	}
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	if err, ok := err.(errs.Error); ok {
		return c.Status(errs.StatusCode(err)).JSON(Error(err))
	}

	log.Printf("caught error on request: %#+v", err)

	return c.SendStatus(http.StatusInternalServerError)
}
