package item

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/pablouser1/NotesManager/db"
)

func Open(myApp fyne.App) {
	itemWindow := myApp.NewWindow("Item")

	// Build form
	entry := widget.NewEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: entry}},
		OnSubmit: func() {
			fmt.Println("Form submitted:", entry.Text)
			db.AddSubject(entry.Text)
			itemWindow.Close()
		},
	}
	itemWindow.SetContent(form)
	itemWindow.Show()
}
