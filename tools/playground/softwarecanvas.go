package playground

import (
	"github.com/emersonkopp/fyne/driver/software"
	"github.com/emersonkopp/fyne/test"
)

// NewSoftwareCanvas creates a new canvas in memory that can render without hardware support
func NewSoftwareCanvas() test.WindowlessCanvas {
	return software.NewCanvas()
}
