package dropbox

//import (
//	"net/http"
//)

const (
	AccessToken string = "redacted"
)

type DropboxStorageClient struct {
}

func (d *DropboxStorageClient) GetFile(path string) ([]byte, error) {
	// curl https://content.dropboxapi.com/2/files/download --header 'Authorization: Bearer <access token>' --header 'Dropbox-API-Arg: {"path":"/<file path>"}'
	return []byte("filecontents"), nil
}

func NewClient() *DropboxStorageClient {
	return &DropboxStorageClient{}
}
