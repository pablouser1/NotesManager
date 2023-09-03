package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"fyne.io/fyne/v2"
	"github.com/pablouser1/NotesManager/constants/editors"
	"github.com/pablouser1/NotesManager/constants/files"
	"github.com/pablouser1/NotesManager/models"
)

func LaunchEditor(myApp fyne.App, subject models.Subject, unit models.Unit) {
	docsPath := myApp.Preferences().String("docs")
	editor := myApp.Preferences().String("editor")
	editorPath := editors.PATHS[editor]
	editorFormat := editors.FORMATS[editor]

	parentPath := filepath.Join(docsPath, subject.Slug)
	os.MkdirAll(parentPath, files.DATA_PERMS)

	path := filepath.Join(parentPath, strconv.FormatInt(unit.Num, 10)+editorFormat)
	// Workaround: If editor is rnote copy the template
	if editor == "rnote" {
		if _, err := os.Stat(path); err != nil {
			if err != nil {
				fmt.Println("Error reading template", err)
				return
			}

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
