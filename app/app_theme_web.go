//go:build !ci && !js && !wasm && test_web_driver
// +build !ci,!js,!wasm,test_web_driver

package app

import (
	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/theme"
)

func defaultVariant() fyne.ThemeVariant {
	return theme.VariantDark
}
