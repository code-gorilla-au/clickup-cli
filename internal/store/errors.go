package store

import "errors"

var (
	// ErrSave - could not save file
	ErrSave = errors.New("error saving")
	// ErrMashaling - could not marshal the data
	ErrMashaling = errors.New("error marshaling data")
	// ErrLoad - error loading
	ErrLoad = errors.New("error loading")
)
