package linkmanager

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/italorfeitosa/go-shorten/pkg/fiberutil"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Container struct {
	FiberApp       *fiber.App
	DB             *gorm.DB
	LinkUseCase    LinkUseCase
	LinkController LinkController
	LinkRepository LinkRepository
}

func Setup() *Container {
	di := new(Container)

	provideDB(di)
	provideFiberApp(di)
	provideLinkRepository(di)
	provideLinkUseCase(di)
	provideLinkController(di)
	provideLinkRouter(di)

	return di
}

func provideDB(di *Container) {
	var err error

	di.DB, err = gorm.Open(sqlite.Open("linkmanager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = di.DB.AutoMigrate(Link{})
	if err != nil {
		log.Fatal(err)
	}
}

func provideFiberApp(di *Container) {
	di.FiberApp = fiber.New(fiber.Config{
		ErrorHandler: fiberutil.ErrorHandler,
	})

	di.FiberApp.Use(recover.New())
	di.FiberApp.Use(logger.New(logger.ConfigDefault))
	di.FiberApp.Use(func(c *fiber.Ctx) error {
		ownerID := c.Get("x-owner-id")
		if ownerID == "" {
			return ErrMissingOwnerID
		}

		c.Locals("owner", ownerID)
		return c.Next()
	})

}

func provideLinkRepository(di *Container) {
	di.LinkRepository = NewLinkRepository(di.DB)
}

func provideLinkUseCase(di *Container) {
	di.LinkUseCase = NewLinkUseCaseHashIDDecorator(NewLinkUseCase(di.LinkRepository))
}

func provideLinkController(di *Container) {
	di.LinkController = NewLinkController(di.LinkUseCase)
}
