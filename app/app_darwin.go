//go:build !ci && !wasm && !test_web_driver && !mobile

package app

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#include <stdbool.h>
#include <stdlib.h>

bool isBundled();
void sendNotification(char *title, char *content);
*/
import "C"
import (
	"fmt"
	"os/exec"
	"strings"
	"unsafe"

	"github.com/emersonkopp/fyne"
)

func (a *fyneApp) SendNotification(n *fyne.Notification) {
	if C.isBundled() {
		titleStr := C.CString(n.Title)
		defer C.free(unsafe.Pointer(titleStr))
		contentStr := C.CString(n.Content)
		defer C.free(unsafe.Pointer(contentStr))

		C.sendNotification(titleStr, contentStr)
		return
	}

	fallbackNotification(n.Title, n.Content)
}

func escapeNotificationString(in string) string {
	noSlash := strings.ReplaceAll(in, "\\", "\\\\")
	return strings.ReplaceAll(noSlash, "\"", "\\\"")
}

//export fallbackSend
func fallbackSend(cTitle, cContent *C.char) {
	title := C.GoString(cTitle)
	content := C.GoString(cContent)
	fallbackNotification(title, content)
}

func fallbackNotification(title, content string) {
	template := `display notification "%s" with title "%s"`
	script := fmt.Sprintf(template, escapeNotificationString(content), escapeNotificationString(title))

	err := exec.Command("osascript", "-e", script).Start()
	if err != nil {
		fyne.LogError("Failed to launch darwin notify script", err)
	}
}
