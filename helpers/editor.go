package helpers

import "fyne.io/fyne/v2"

func GetEditor(myApp fyne.App) string {
	return myApp.Preferences().String("editor")
}

func SetEditor(myApp fyne.App, val string) {
	myApp.Preferences().SetString("editor", val)
}
