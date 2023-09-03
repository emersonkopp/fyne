package mobile

import (
	"github.com/emersonkopp/fyne"
)

// Declare conformity with Clipboard interface
var _ fyne.Clipboard = (*mobileClipboard)(nil)

// mobileClipboard represents the system mobileClipboard
type mobileClipboard struct {
}
