package ui

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)

	app.AppWindow.SetContent(swatchesContainer)
}
