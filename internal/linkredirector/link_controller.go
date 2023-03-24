package linkredirector

import (
	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-shorten/pkg/hashid"
)

type LinkController interface {
	RedirectTo(c *fiber.Ctx) error
}

type linkController struct {
	linkUseCase LinkUseCase
}

func (ctl *linkController) RedirectTo(c *fiber.Ctx) error {
	hash := hashid.FromString(c.Params("hash"))

	originalLink, err := ctl.linkUseCase.GetOriginalLink(c.UserContext(), hash.ID())
	if err != nil {
		return err
	}

	return c.Redirect(originalLink)
}
