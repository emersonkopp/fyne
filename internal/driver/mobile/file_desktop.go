//go:build !ios && !android
// +build !ios,!android

package mobile

import (
	"io"

	"github.com/emersonkopp/fyne"
	intRepo "github.com/emersonkopp/fyne/internal/repository"
	"github.com/emersonkopp/fyne/storage/repository"
)

func existsURI(fyne.URI) (bool, error) {
	// no-op as we use the internal FileRepository
	return false, nil
}

func nativeFileOpen(*fileOpen) (io.ReadCloser, error) {
	// no-op as we use the internal FileRepository
	return nil, nil
}

func nativeFileSave(*fileSave) (io.WriteCloser, error) {
	// no-op as we use the internal FileRepository
	return nil, nil
}

func registerRepository(d *mobileDriver) {
	repository.Register("file", intRepo.NewFileRepository())
}
