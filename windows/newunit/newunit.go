package newunit

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/pablouser1/NotesManager/constants/ui"
	"github.com/pablouser1/NotesManager/db"
	"github.com/pablouser1/NotesManager/models"
)

func Open(myApp fyne.App, subject models.Subject, channel chan models.Unit) {
	itemWindow := myApp.NewWindow("New Unit")
	itemWindow.Resize(ui.MISC_WIN_SIZE)

	itemWindow.SetCloseIntercept(func() {
		channel <- models.Unit{}
		itemWindow.Close()
	})

	// Build form
	num := widget.NewEntry()
	name := widget.NewEntry()
	subjectEntry := widget.NewEntry()
	subjectEntry.SetText(subject.Name)
	subjectEntry.Disable()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Unit number",
				Widget: num,
			},
			{
				Text:   "Unit name",
				Widget: name,
			},
			{
				Text:   "Bind to subject",
				Widget: subjectEntry,
			},
		},
		OnSubmit: func() {
			fmt.Println("Form submitted:", num.Text, name.Text, subjectEntry.Text)
			realNumInt, err := strconv.Atoi(num.Text)
			if err != nil {
				fmt.Println("Error sending form", err)
			}

			realNum := int64(realNumInt)

			unit, err := db.AddUnit(realNum, name.Text, subject.ID)
			if err != nil {
				fmt.Println("Error writing unit to db", err)
			}

			channel <- unit

			itemWindow.Close()
		},
	}
	itemWindow.SetContent(form)
	itemWindow.Show()
}
