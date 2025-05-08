package app_test

import (
	"testing"

	"github.com/emersonkopp/fyne/internal/app"
	"github.com/emersonkopp/fyne/test"
)

func TestApplySettings_BeforeContentSet(t *testing.T) {
	a := test.NewApp()
	w := a.NewWindow("NoContent")
	defer w.Close()

	app.ApplySettings(a.Settings(), a)
}
