package test

import (
	"os"

	"github.com/emersonkopp/fyne"
	"github.com/emersonkopp/fyne/internal"
	"github.com/emersonkopp/fyne/storage"
)

type testStorage struct {
	*internal.Docs
}

func (s *testStorage) RootURI() fyne.URI {
	return storage.NewFileURI(os.TempDir())
}

func (s *testStorage) docRootURI() (fyne.URI, error) {
	return storage.Child(s.RootURI(), "Documents")
}
