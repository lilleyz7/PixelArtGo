package pixl

import (
	"image/color"
	"pxlart/appTypes"
	"pxlart/swatch"
	"pxlart/ui"

	"fyne.io/fyne/v2/app"
)

func Run() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := appTypes.AppState{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		AppWindow: pixlWindow,
		State:     &state,
		Swatches:  make([]*swatch.Swatch, 0, 64),
	}
	ui.Setup(&appInit)

	appInit.AppWindow.ShowAndRun()

}
