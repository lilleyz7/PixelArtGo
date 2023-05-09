package appTypes

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type CanvasConfig struct {
	DrawArea       fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}

type AppState struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int
	FilePath       string
}

func (state *AppState) SetFilePath(path string) {
	state.FilePath = path
}
