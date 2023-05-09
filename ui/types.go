package ui

import (
	"pxlart/appTypes"
	"pxlart/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	AppWindow fyne.Window
	State     *appTypes.AppState
	Swatches  []*swatch.Swatch
}
