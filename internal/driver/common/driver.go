package common

import (
	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/internal/cache"
)

// CanvasForObject returns the canvas for the specified object.
func CanvasForObject(obj fyne.CanvasObject) fyne.Canvas {
	return cache.GetCanvasForObject(obj)
}
