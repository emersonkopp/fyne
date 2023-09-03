//go:build !ci && (!android || !ios || !mobile) && (js || wasm || test_web_driver)
// +build !ci
// +build !android !ios !mobile
// +build js wasm test_web_driver

package app

import (
	"github.com/emersonkopp/fyne"
)

func (app *fyneApp) SendNotification(_ *fyne.Notification) {
	// TODO #2735
	fyne.LogError("Sending notification is not supported yet.", nil)
}

func rootConfigDir() string {
	return "/data/"
}
