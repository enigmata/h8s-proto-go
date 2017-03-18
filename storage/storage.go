package storage

import (
	"godel/device"
	"godel/storage/dropbox"
)

type StorageClient interface {
	GetFile(string) ([]byte, error)
}

const (
	DROPBOX string = "dropbox"
)

func NewClient(name string) StorageClient {
	switch name {
	case DROPBOX:
		return dropbox.NewClient()
	default:
		return nil
	}
	return nil
}
