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
	"github.com/pablouser1/NotesManager/helpers/dav"
	"github.com/pablouser1/NotesManager/models"
)

func getEditor(myApp fyne.App) models.Editor {
	editor := myApp.Preferences().String("editor")
	format := editors.FORMATS[editor]
	var path string
	switch runtime.GOOS {
	case "linux":
		path = editors.PATHS_LINUX[editor]
	case "windows":
		path = editors.PATHS_WINDOWS[editor]
	default:
		path = ""
	}

	return models.Editor{
		Name:   editor,
		Format: format,
		Path:   path,
	}
}

func buildFilename(unit models.Unit, variant models.Variant, format string) string {
	filename := strconv.FormatInt(unit.Num, 10)

	if variant != (models.Variant{}) {
		filename += "-" + variant.Slug
	}

	filename += format

	return filename
}

func LaunchEditor(myApp fyne.App, subject models.Subject, unit models.Unit, variant models.Variant) {
	docsPath := myApp.Preferences().String("docs")
	editor := getEditor(myApp)

	filename := buildFilename(unit, variant, editor.Format)

	relPath := filepath.Join(subject.Slug, filename)

	parentPath := filepath.Join(docsPath, subject.Slug)
	os.MkdirAll(parentPath, files.DATA_PERMS)

	path := filepath.Join(docsPath, relPath)
	// Workaround: If editor is rnote copy the template
	if editor.Name == "rnote" {
		if _, err := os.Stat(path); err != nil {
			err = os.WriteFile(path, files.RNOTE_TEMPLATE, files.DATA_PERMS)
			if err != nil {
				fmt.Println("Error writing template", err)
				return
			}
		}
	}

	cmd := exec.Command(editor.Path, path)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error when executing editor", err)
		return
	}

	// Sync with WebDav
	if dav.IsEnabled(myApp) {
		go func() {
			dav.NewClient(dav.GetConfig(myApp), docsPath)
			err := dav.Upload(relPath)
			if err == nil {
				myApp.SendNotification(
					fyne.NewNotification(
						"NotesManager",
						fmt.Sprintf("%s_%d uploaded successfully", subject.Slug, unit.Num),
					),
				)
			} else {
				myApp.SendNotification(
					fyne.NewNotification(
						"NotesManager",
						fmt.Sprintf("Error uploading %s-%d: %e", subject.Slug, unit.Num, err),
					),
				)
			}
		}()
	}
}
