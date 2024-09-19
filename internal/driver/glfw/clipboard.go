//go:build !wasm && !test_web_driver

package glfw

import (
	"runtime"
	"time"

	"github.com/emersonkopp/fyne"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// Declare conformity with Clipboard interface
var _ fyne.Clipboard = (*clipboard)(nil)

// clipboard represents the system clipboard
type clipboard struct{}

// Content returns the clipboard content
func (c *clipboard) Content() string {
	// This retry logic is to work around the "Access Denied" error often thrown in windows PR#1679
	if runtime.GOOS != "windows" {
		return c.content()
	}
	for i := 3; i > 0; i-- {
		cb := c.content()
		if cb != "" {
			return cb
		}
		time.Sleep(50 * time.Millisecond)
	}
	// can't log retry as it would alos log errors for an empty clipboard
	return ""
}

func (c *clipboard) content() string {
	content := ""
	runOnMain(func() {
		content = glfw.GetClipboardString()
	})
	return content
}

// SetContent sets the clipboard content
func (c *clipboard) SetContent(content string) {
	// This retry logic is to work around the "Access Denied" error often thrown in windows PR#1679
	if runtime.GOOS != "windows" {
		c.setContent(content)
		return
	}
	for i := 3; i > 0; i-- {
		c.setContent(content)
		if c.content() == content {
			return
		}
		time.Sleep(50 * time.Millisecond)
	}
	fyne.LogError("GLFW clipboard set failed", nil)
}

func (c *clipboard) setContent(content string) {
	runOnMain(func() {
		glfw.SetClipboardString(content)
	})
}
