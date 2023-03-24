package linkmanager

func provideLinkRouter(di *Container) {
	di.FiberApp.Group("/linkmanager").
		Post("/links", di.LinkController.Create).
		Get("/links", di.LinkController.GetAll).
		Patch("/links/:hash/info", di.LinkController.UpdateInfo).
		Patch("/links/:hash/originallink", di.LinkController.ChangeOriginaLink).
		Delete("/links/:hash", di.LinkController.Delete)
}
