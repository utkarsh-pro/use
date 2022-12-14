package storage

import (
	"fmt"
	"io"

	"github.com/utkarsh-pro/use/pkg/storage/config"
	"github.com/utkarsh-pro/use/pkg/storage/stupid"
)

type Storage interface {
	// Init configures the storage.
	Init() error

	// Get returns the value for the given key.
	Get(key []byte) ([]byte, error)

	// Set sets the value for the given key.
	Set(key []byte, value []byte) error

	// Delete deletes the value for the given key.
	Delete(key []byte) error

	// Exists returns true if the given key exists.
	Exists(key []byte) (bool, error)

	// Len returns the number of keys in the storage.
	Len() (int, error)

	// PhysicalSnapshot writes snapshot of the storage data to
	// the given writer.
	PhysicalSnapshot(w io.Writer) error

	// Close closes the storage.
	Close() error
}

type StorageType string

const (
	StupidStorageType StorageType = "stupid"
)

// New returns a new Storage instance.
func New(t StorageType, path string, config config.Config) (Storage, error) {
	switch t {
	case StupidStorageType:
		storage := stupid.New(path, config)
		return storage, nil
	default:
		return nil, fmt.Errorf("unknown storage type: %s", t)
	}
}
