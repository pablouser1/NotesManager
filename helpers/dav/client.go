package dav

import (
	"os"

	"github.com/pablouser1/NotesManager/constants/files"
	"github.com/studio-b12/gowebdav"
)

var c *gowebdav.Client

func Connect(host string, username string, password string, base string) error {
	c = gowebdav.NewClient(host, username, password)
	err := c.Connect()
	// Create base path
	if err == nil {
		c.MkdirAll(base, files.DATA_PERMS)
	}
	return err
}

func Upload(localPath string, remotePath string) {
	file, _ := os.Open(localPath)
	defer file.Close()

	c.WriteStream("", file, files.DATA_PERMS)
}
