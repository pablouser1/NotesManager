package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"fyne.io/fyne/v2"
	"github.com/pablouser1/NotesManager/constants/editors"
	"github.com/pablouser1/NotesManager/constants/files"
	"github.com/pablouser1/NotesManager/models"
)

func LaunchEditor(myApp fyne.App, subject models.Subject, unit models.Unit) {
	docsPath := myApp.Preferences().String("docs")
	editor := myApp.Preferences().String("editor")
	editorFormat := editors.FORMATS[editor]

	var editorPath string
	switch runtime.GOOS {
	case "linux":
		editorPath = editors.PATHS_LINUX[editor]
	case "windows":
		editorPath = editors.PATHS_WINDOWS[editor]
	default:
		fmt.Println("OS Not supported")
		return
	}

	parentPath := filepath.Join(docsPath, subject.Slug)
	os.MkdirAll(parentPath, files.DATA_PERMS)

	path := filepath.Join(parentPath, strconv.FormatInt(unit.Num, 10)+editorFormat)
	// Workaround: If editor is rnote copy the template
	if editor == "rnote" {
		if _, err := os.Stat(path); err != nil {
			err = os.WriteFile(path, files.RNOTE_TEMPLATE, files.DATA_PERMS)
			if err != nil {
				fmt.Println("Error writing template", err)
				return
			}
		}
	}

	cmd := exec.Command(editorPath, path)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error when executing editor", err)
		return
	}
}
