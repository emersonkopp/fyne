//go:build linux || freebsd || openbsd || netbsd

package glfw

import "github.com/emersonkopp/fyne"

func (w *window) platformResize(canvasSize fyne.Size) {
	w.canvas.Resize(canvasSize)
}
