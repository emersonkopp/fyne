//go:build !hints

package theme

import (
	"image/color"

	"github.com/emersonkopp/fyne"
)

var (
	fallbackColor = color.Transparent
	fallbackIcon  = &fyne.StaticResource{}
)
