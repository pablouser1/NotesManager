package dav

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pablouser1/NotesManager/constants/files"
	"github.com/pablouser1/NotesManager/models"
	"github.com/studio-b12/gowebdav"
)

var c *gowebdav.Client
var connected bool = false
var davBase string
var localBase string

func NewClient(webdav models.WebDav, localPath string) error {
	c = gowebdav.NewClient(webdav.Host, webdav.Username, webdav.Password)
	c.SetInterceptor(func(method string, rq *http.Request) {
		// Only for uploading
		// Taken from https://github.com/studio-b12/gowebdav/issues/35#issuecomment-827806721
		if rq.Method == "PUT" {
			b, err := io.ReadAll(rq.Body)
			if err != nil {
				panic(err)
			}

			rq.ContentLength = int64(len(b))

			rq.Body = io.NopCloser(bytes.NewReader(b))
		}
	})
	err := c.Connect()

	connected = err == nil
	// Create base path
	if connected {
		c.MkdirAll(webdav.Base, files.DATA_PERMS)
	}

	davBase = webdav.Base
	localBase = localPath

	return err
}

func Upload(relPath string) error {
	if !connected {
		return errors.New("not connected to webdav")
	}

	absLocalPath := filepath.Join(localBase, relPath)
	absDavPath := filepath.Join(davBase, relPath)

	file, err := os.Open(absLocalPath)

	if err != nil {
		return err
	}

	err = c.WriteStream(absDavPath, file, files.DATA_PERMS)
	return err
}
