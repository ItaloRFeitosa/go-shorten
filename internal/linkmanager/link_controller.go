package linkmanager

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/italorfeitosa/go-shorten/pkg/fiberutil"
	"github.com/italorfeitosa/go-shorten/pkg/hashid"
	"github.com/italorfeitosa/go-shorten/pkg/validate"
)

type LinkController interface {
	Create(c *fiber.Ctx) error
	UpdateInfo(c *fiber.Ctx) error
	ChangeOriginaLink(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

type linkController struct {
	linkUseCase LinkUseCase
}

func NewLinkController(u LinkUseCase) *linkController {
	return &linkController{u}
}

func (ctl *linkController) Create(c *fiber.Ctx) error {
	var input CreateLinkInput

	input.OwnerID = c.Locals("owner").(string)

	if err := fiberutil.Body(c, &input); err != nil {
		return err
	}

	if err := ctl.linkUseCase.Create(c.UserContext(), input); err != nil {
		return err
	}

	return c.SendStatus(http.StatusCreated)
}

func (ctl *linkController) GetAll(c *fiber.Ctx) error {
	var query GetAllLinksQuery

	query.OwnerID = c.Locals("owner").(string)

	if err := fiberutil.Query(c, &query); err != nil {
		return err
	}

	if query.Limit <= 0 || query.Limit > 100 {
		query.Limit = 100
	}

	if query.Page <= 0 {
		query.Page = 1
	}

	links, err := ctl.linkUseCase.Find(c.UserContext(), query)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiberutil.Data(links))
}

func (ctl *linkController) UpdateInfo(c *fiber.Ctx) error {
	var input UpdateLinkInfoInput

	input.OwnerID = c.Locals("owner").(string)
	input.Hash = hashid.FromString(c.Params("hash"))

	if err := fiberutil.Body(c, &input); err != nil {
		return err
	}

	if err := ctl.linkUseCase.UpdateInfo(c.UserContext(), input); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

func (ctl *linkController) ChangeOriginaLink(c *fiber.Ctx) error {
	var input ChangeOriginalLinkInput

	input.OwnerID = c.Locals("owner").(string)
	input.Hash = hashid.FromString(c.Params("hash"))

	if err := fiberutil.Body(c, &input); err != nil {
		return err
	}

	if err := ctl.linkUseCase.ChangeOriginalLink(c.UserContext(), input); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}

func (ctl *linkController) Delete(c *fiber.Ctx) error {
	var input DeleteLinkInput

	input.OwnerID = c.Locals("owner").(string)
	input.Hash = hashid.FromString(c.Params("hash"))

	if err := validate.Struct(input); err != nil {
		return err
	}

	if err := ctl.linkUseCase.Delete(c.UserContext(), input); err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)
}
