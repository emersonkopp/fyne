package dialog

import (
	"testing"

	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/test"
)

func Test_colorChannel_Layout(t *testing.T) {
	test.NewTempApp(t)

	min := 0
	max := 100
	size := fyne.NewSize(250, 50)

	for name, tt := range map[string]struct {
		name  string
		value int
	}{
		"foobar_0": {
			name:  "foobar",
			value: 0,
		},
		"foobar_50": {
			name:  "foobar",
			value: 50,
		},
		"foobar_100": {
			name:  "foobar",
			value: 100,
		},
	} {
		t.Run(name, func(t *testing.T) {
			color := newColorChannel(tt.name, min, max, tt.value, nil)
			color.Resize(size)

			window := test.NewTempWindow(t, color)

			test.AssertRendersToImage(t, "color/channel_layout_"+name+".png", window.Canvas())
		})
	}
}
