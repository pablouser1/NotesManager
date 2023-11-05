package variant

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/pablouser1/NotesManager/db"
	"github.com/pablouser1/NotesManager/helpers"
	"github.com/pablouser1/NotesManager/models"
)

func getVariantsName() []string {
	variants, err := db.GetVariants()

	if err != nil {
		fmt.Println("GetVariants:", err)
		return []string{}
	}

	var variantsRadio []string

	for _, s := range variants {
		variantsRadio = append(variantsRadio, s.Name)
	}

	return variantsRadio
}

func runEditor(myApp fyne.App, subject models.Subject, unit models.Unit, variant models.Variant, mainWindow fyne.Window) {
	mainWindow.Hide()
	go func() {
		helpers.LaunchEditor(myApp, subject, unit, variant)
		mainWindow.Show()
	}()
}

func NewVariantDialog(myApp fyne.App, mainWindow fyne.Window, subject models.Subject, unit models.Unit) dialog.Dialog {
	variantsName := getVariantsName()

	choiseValue := variantsName[0]

	choise := widget.NewSelect(variantsName, func(s string) {
		choiseValue = s
	})
	choise.SetSelectedIndex(0)

	if len(variantsName) == 0 {
		choise.Disable()
	}

	variantDialog := dialog.NewCustomConfirm(subject.Name+" :: "+strconv.FormatInt(unit.Num, 10), "Go", "Cancel", choise, func(b bool) {
		if !b {
			return
		}

		var variant models.Variant
		var err error

		if choiseValue != "" {
			variant, err = db.GetVariantByName(choiseValue)
			if err != nil {
				fmt.Println("GetVariantByName:", err)
				return
			}
		}

		runEditor(myApp, subject, unit, variant, mainWindow)
	}, mainWindow)

	return variantDialog
}
