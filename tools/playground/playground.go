// Package playground provides tooling for running fyne applications inside the Go playground.
package playground

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"

	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/driver/software"
	"github.com/emersonkopp/fyne/internal/test"
	"github.com/emersonkopp/fyne/theme"
)

func imageToPlayground(img image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		fyne.LogError("Failed to encode image", err)
		return
	}

	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Println("IMAGE:" + enc)
}

// RenderCanvas takes a canvas and converts it into an inline image for showing in the playground
func RenderCanvas(c fyne.Canvas) {
	imageToPlayground(software.RenderCanvas(c, test.DarkTheme(theme.DefaultTheme())))
}

// RenderWindow takes a window and converts it's canvas into an inline image for showing in the playground
func RenderWindow(w fyne.Window) {
	RenderCanvas(w.Canvas())
}

// Render takes a canvasobject and converts it into an inline image for showing in the playground
func Render(obj fyne.CanvasObject) {
	imageToPlayground(software.Render(obj, test.DarkTheme(theme.DefaultTheme())))
}
