package main

import (
	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/app"
	"github.com/emersonkopp/fyne/cmd/fyne_settings/settings"
	"github.com/emersonkopp/fyne/container"
)

func main() {
	s := settings.NewSettings()

	a := app.New()
	w := a.NewWindow("Fyne Settings")

	appearance := s.LoadAppearanceScreen(w)
	tabs := container.NewAppTabs(
		&container.TabItem{Text: "Appearance", Icon: s.AppearanceIcon(), Content: appearance})
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)

	w.Resize(fyne.NewSize(520, 520))
	w.ShowAndRun()
}
