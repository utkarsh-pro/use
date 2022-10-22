package storage

import (
	"fmt"

	"github.com/utkarsh-pro/use/pkg/storage/stupid"
)

type Storage interface {
	// Init configures the storage.
	Init() error

	// Get returns the value for the given key.
	Get(key string) ([]byte, error)

	// Set sets the value for the given key.
	Set(key string, value []byte) error

	// Delete deletes the value for the given key.
	Delete(key string) error

	// Exists returns true if the given key exists.
	Exists(key string) (bool, error)

	// Len returns the number of keys in the storage.
	Len() (int, error)

	// Close closes the storage.
	Close() error
}

type StorageType string

const (
	StupidStorageType StorageType = "stupid"
)

// New returns a new Storage instance.
func New(t StorageType, path string) (Storage, error) {
	switch t {
	case StupidStorageType:
		storage := stupid.New(path)
		return storage, nil
	default:
		return nil, fmt.Errorf("unknown storage type: %s", t)
	}
}
