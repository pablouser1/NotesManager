package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/pablouser1/NotesManager/constants/editors"
	"github.com/pablouser1/NotesManager/constants/ui"
	"github.com/pablouser1/NotesManager/helpers"
	"github.com/pablouser1/NotesManager/helpers/dav"
)

func buildGeneral(myApp fyne.App, settingsWindow fyne.Window, submitDialog dialog.Dialog) fyne.Widget {
	docs := widget.NewEntry()

	// Build file opener
	fo := dialog.NewFolderOpen(func(uc fyne.ListableURI, err error) {
		// File opener callback
		if err != nil || uc == nil {
			return
		}

		docs.SetText(uc.Path())
	}, settingsWindow)

	// Setup entry
	docs.ActionItem = widget.NewButtonWithIcon("", theme.FolderOpenIcon(), fo.Show)
	docs.SetText(myApp.Preferences().String("docs"))

	// Editor picker
	editor := helpers.GetEditor(myApp)
	editors := widget.NewSelect(editors.AVAILABLE, func(value string) {
		editor = value
	})
	editors.SetSelected(editor)

	// Build form
	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Docs path",
				Widget: docs,
			},
			{
				Text:   "Editor",
				Widget: editors,
			},
		},
		OnSubmit: func() {
			helpers.SetDocs(myApp, docs.Text)
			helpers.SetEditor(myApp, editor)
			submitDialog.Show()
		},
	}

	return form
}

func buildWebDAV(myApp fyne.App, settingsWindow fyne.Window, submitDialog dialog.Dialog) fyne.Widget {
	webdav := dav.GetConfig(myApp)

	// Enabled
	enabled := widget.NewCheck("Enabled", func(b bool) {
		webdav.Enabled = b
	})
	enabled.SetChecked(webdav.Enabled)

	// Host
	host := widget.NewEntry()
	host.SetText(webdav.Host)

	// Username
	username := widget.NewEntry()
	username.SetText(webdav.Username)

	// Password
	password := widget.NewEntry()
	password.SetText(webdav.Password)

	// Base
	base := widget.NewEntry()
	base.SetText(webdav.Base)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   "Status",
				Widget: enabled,
			},
			{
				Text:   "Host",
				Widget: host,
			},
			{
				Text:   "Username",
				Widget: username,
			},
			{
				Text:   "Password",
				Widget: password,
			},
			{
				Text:   "Base",
				Widget: base,
			},
		},
		OnSubmit: func() {
			webdav.Host = host.Text
			webdav.Username = username.Text
			webdav.Password = password.Text
			webdav.Base = base.Text
			dav.SetConfig(myApp, webdav)
			submitDialog.Show()
		},
	}

	return form
}

func Open(myApp fyne.App) {
	settingsWindow := myApp.NewWindow("Settings")
	settingsWindow.Resize(ui.MISC_WIN_SIZE)

	submitDialog := dialog.NewInformation("OK", "Settings saved", settingsWindow)

	generalForm := buildGeneral(myApp, settingsWindow, submitDialog)
	davForm := buildWebDAV(myApp, settingsWindow, submitDialog)

	tabs := container.NewAppTabs(
		container.NewTabItem("General", generalForm),
		container.NewTabItem("WebDAV", davForm),
	)

	settingsWindow.SetContent(tabs)
	settingsWindow.Show()
}
