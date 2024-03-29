package main

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/pablouser1/NotesManager/constants/files"
	"github.com/pablouser1/NotesManager/db"
	"github.com/pablouser1/NotesManager/helpers"
	"github.com/pablouser1/NotesManager/windows/home"
)

func main() {
	myApp := app.NewWithID("es.pablouser1.notesmanager")
	configPath := myApp.Storage().RootURI().Path()

	// Create default path for docs and write initial config
	docsPath := helpers.GetDocs(myApp)
	if docsPath == "" {
		defaultPath := filepath.Join(configPath, "docs")
		err := os.MkdirAll(defaultPath, files.DATA_PERMS)
		if err != nil {
			fmt.Println(err)
			return
		}
		helpers.SetDocs(myApp, defaultPath)
	}

	// Set default editor if empty
	editor := helpers.GetEditor(myApp)
	if editor == "" {
		helpers.SetEditor(myApp, "rnote")
	}

	// System tray
	if desk, ok := myApp.(desktop.App); ok {
		m := fyne.NewMenu("Notes Manager",
			fyne.NewMenuItem("Show", home.Show),
		)
		desk.SetSystemTrayMenu(m)
	}

	// Get ready DB
	db.Open(configPath)

	// Open main menu
	home.Open(myApp)

	// Run app
	myApp.Run()

	// Cleanup
	db.Close()
}
