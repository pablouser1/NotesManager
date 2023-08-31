package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
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
		err := os.MkdirAll(configPath+"/docs", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
		helpers.SetDocs(myApp, docsPath)
	}

	// Set default editor if empty
	editor := helpers.GetEditor(myApp)
	if editor == "" {
		helpers.SetEditor(myApp, "rnote")
	}

	// System tray
	if desk, ok := myApp.(desktop.App); ok {
		m := fyne.NewMenu("Notes Manager",
			fyne.NewMenuItem("Show", func() {
				home.Show()
			}),
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
