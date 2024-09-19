package dialog

import (
	"image/color"
	"testing"

	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/test"
	"github.com/emersonkopp/fyne/theme"
)

func Test_colorPreview_Color(t *testing.T) {
	test.NewTempApp(t)

	preview := newColorPreview(color.RGBA{53, 113, 233, 255})
	preview.SetColor(color.RGBA{90, 206, 80, 180})
	window := test.NewTempWindow(t, preview)
	padding := theme.Padding() * 2
	window.Resize(fyne.NewSize(128+padding, 64+padding))

	test.AssertRendersToImage(t, "color/preview_color.png", window.Canvas())
}
