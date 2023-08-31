package helpers

import "fyne.io/fyne/v2"

func GetDocs(myApp fyne.App) string {
	return myApp.Preferences().String("docs")
}

func SetDocs(myApp fyne.App, val string) {
	myApp.Preferences().SetString("docs", val)
}
