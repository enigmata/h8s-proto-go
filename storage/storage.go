package storage

import (
	//"github.com/enigmata/h8s-proto-go/device"
	"github.com/enigmata/h8s-proto-go/storage/dropbox"
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
