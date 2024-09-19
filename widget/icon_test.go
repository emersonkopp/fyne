package widget_test

import (
	"testing"

	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/layout"
	"github.com/emersonkopp/fyne/test"
	"github.com/emersonkopp/fyne/theme"
	"github.com/emersonkopp/fyne/widget"
)

func TestIcon_Layout(t *testing.T) {
	test.NewTempApp(t)

	for name, tt := range map[string]struct {
		resource fyne.Resource
	}{
		"empty": {},
		"resource": {
			resource: theme.CancelIcon(),
		},
	} {
		t.Run(name, func(t *testing.T) {
			icon := &widget.Icon{
				Resource: tt.resource,
			}

			window := test.NewTempWindow(t, &fyne.Container{Layout: layout.NewCenterLayout(), Objects: []fyne.CanvasObject{icon}})
			window.Resize(icon.MinSize().Max(fyne.NewSize(150, 200)))

			test.AssertRendersToMarkup(t, "icon/layout_"+name+".xml", window.Canvas())
		})
	}
}
