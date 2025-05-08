//go:build !debug && !release

package build

import "github.com/emersonkopp/fyne"

// Mode is the application's build mode.
const Mode = fyne.BuildStandard
