package widget

import (
	"testing"

	"github.com/emersonkopp/fyne/canvas"
	"github.com/emersonkopp/fyne/internal/cache"
	"github.com/emersonkopp/fyne/theme"
	"github.com/stretchr/testify/assert"
)

type extendedIcon struct {
	Icon
}

func newExtendedIcon() *extendedIcon {
	icon := &extendedIcon{}
	icon.ExtendBaseWidget(icon)
	return icon
}

func TestIcon_Extended_SetResource(t *testing.T) {
	icon := newExtendedIcon()
	icon.SetResource(theme.ComputerIcon())

	objs := cache.Renderer(icon).Objects()
	assert.Equal(t, 1, len(objs))
	assert.Equal(t, theme.ComputerIcon(), objs[0].(*canvas.Image).Resource)
}
