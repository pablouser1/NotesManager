package dav

import (
	"fyne.io/fyne/v2"
	"github.com/pablouser1/NotesManager/models"
)

const enabled_key = "webdav_enabled"
const host_key = "webdav_host"
const username_key = "webdav_username"
const password_key = "webdav_password"
const base_key = "webdav_base"

func GetConfig(myApp fyne.App) models.WebDav {
	prefs := myApp.Preferences()
	enabled := IsEnabled(myApp)
	host := prefs.StringWithFallback(host_key, "")
	username := prefs.StringWithFallback(username_key, "")
	password := prefs.StringWithFallback(password_key, "")
	base := prefs.StringWithFallback(base_key, "/NotesManager")

	return models.WebDav{
		Enabled:  enabled,
		Host:     host,
		Username: username,
		Password: password,
		Base:     base,
	}
}

func SetConfig(myApp fyne.App, dav models.WebDav) {
	prefs := myApp.Preferences()
	prefs.SetBool(enabled_key, dav.Enabled)
	prefs.SetString(host_key, dav.Host)
	prefs.SetString(username_key, dav.Username)
	prefs.SetString(password_key, dav.Password)
	prefs.SetString(base_key, dav.Base)
}

func IsEnabled(myApp fyne.App) bool {
	return myApp.Preferences().BoolWithFallback(enabled_key, false)
}
