package storage

import (
	"github.com/emersonkopp/fyne/storage/repository"
)

// URIRootError is a wrapper for repository.URIRootError
//
// Deprecated - use repository.ErrURIRoot instead
var URIRootError = repository.ErrURIRoot
