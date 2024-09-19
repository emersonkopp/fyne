//go:build !wasm && test_web_driver

package app

import (
	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/theme"
)

// DefaultVariant returns the systems default fyne.ThemeVariant.
// Normally, you should not need this. It is extracted out of the root app package to give the
// settings app access to it.
func DefaultVariant() fyne.ThemeVariant {
	return theme.VariantDark
}
