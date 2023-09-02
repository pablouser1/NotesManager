package newsub

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/pablouser1/NotesManager/constants/ui"
	"github.com/pablouser1/NotesManager/db"
)

func Open(myApp fyne.App) {
	itemWindow := myApp.NewWindow("New Subject")
	itemWindow.Resize(ui.MISC_WIN_SIZE)

	// Build form
	entry := widget.NewEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: entry}},
		OnSubmit: func() {
			fmt.Println("Form submitted:", entry.Text)
			_, err := db.AddSubject(entry.Text)
			if err != nil {
				fmt.Println("Error writing subject to db", err)
			}
			itemWindow.Close()
		},
	}
	itemWindow.SetContent(form)
	itemWindow.Show()
}
