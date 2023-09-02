package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/pablouser1/NotesManager/constants/editors"
	"github.com/pablouser1/NotesManager/constants/ui"
	"github.com/pablouser1/NotesManager/helpers"
)

func Open(myApp fyne.App) {
	settingsWindow := myApp.NewWindow("Settings")
	settingsWindow.Resize(ui.MISC_WIN_SIZE)
	entry := widget.NewEntry()

	// Build file opener
	fo := dialog.NewFolderOpen(func(uc fyne.ListableURI, err error) {
		// File opener callback
		if err != nil || uc == nil {
			return
		}

		entry.SetText(uc.Path())
	}, settingsWindow)

	// Setup entry
	entry.ActionItem = widget.NewButtonWithIcon("", theme.FolderOpenIcon(), fo.Show)
	entry.SetText(myApp.Preferences().String("docs"))

	// Editor picker
	editor := helpers.GetEditor(myApp)
	combo := widget.NewSelect(editors.AVAILABLE, func(value string) {
		editor = value
	})
	combo.SetSelected(editor)

	// Build form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Docs path",
				Widget: entry,
			},
			{
				Text:   "Editor",
				Widget: combo,
			},
		},
		OnSubmit: func() {
			helpers.SetDocs(myApp, entry.Text)
			helpers.SetEditor(myApp, editor)
			settingsWindow.Close()
		},
	}
	settingsWindow.SetContent(form)
	settingsWindow.Show()
}
