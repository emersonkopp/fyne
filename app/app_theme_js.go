//go:build !ci && js && !wasm
// +build !ci,js,!wasm

package app

import (
	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/theme"

	"github.com/gopherjs/gopherjs/js"
)

func defaultVariant() fyne.ThemeVariant {
	if matchMedia := js.Global.Call("matchMedia", "(prefers-color-scheme: dark)"); matchMedia != js.Undefined {
		if matches := matchMedia.Get("matches"); matches != js.Undefined && matches.Bool() {
			return theme.VariantDark
		}
		return theme.VariantLight
	}
	return theme.VariantDark
}

func init() {
	if matchMedia := js.Global.Call("matchMedia", "(prefers-color-scheme: dark)"); matchMedia != js.Undefined {
		matchMedia.Call("addEventListener", "change", func(o *js.Object) {
			fyne.CurrentApp().Settings().(*settings).setupTheme()
		})
	}
}
