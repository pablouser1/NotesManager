package helpers

import (
	"fmt"
	"os/exec"
	"strconv"

	"fyne.io/fyne/v2"
	"github.com/pablouser1/NotesManager/constants/editors"
	"github.com/pablouser1/NotesManager/models"
)

func LaunchEditor(myApp fyne.App, subject models.Subject, unit models.Unit) {
	docsPath := myApp.Preferences().String("docs")
	editor := myApp.Preferences().String("editor")
	editorPath := editors.PATHS[editor]
	editorFormat := editors.FORMATS[editor]

	path := docsPath + "/" + subject.Slug + "/" + strconv.FormatInt(unit.Num, 10) + editorFormat
	cmd := exec.Command(editorPath, path)

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
