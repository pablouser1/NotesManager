package home

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/pablouser1/NotesManager/constants/ui"
	"github.com/pablouser1/NotesManager/db"
	"github.com/pablouser1/NotesManager/helpers"
	"github.com/pablouser1/NotesManager/models"
	"github.com/pablouser1/NotesManager/windows/newsub"
	"github.com/pablouser1/NotesManager/windows/newunit"
	"github.com/pablouser1/NotesManager/windows/settings"
)

var mainWindow fyne.Window

func getIndexesFromList[T any](arr []T) []int {
	var res []int
	for i := range arr {
		res = append(res, i)
	}

	return res
}

func Open(myApp fyne.App) {
	mainWindow = myApp.NewWindow("Notes Manager")
	mainWindow.Resize(ui.MAIN_WIN_SIZE)
	mainWindow.SetCloseIntercept(func() {
		mainWindow.Hide()
	})

	// Picked subject
	var subject models.Subject

	// Get subjects
	subjects, err := db.GetSubjects()
	if err != nil {
		fmt.Println("GetSubjects:", err)
		return
	}

	// Set default selected subject
	subject = models.Subject{}
	if len(subjects) > 0 {
		subject = subjects[0]
	}

	// Get default units
	units, err := db.GetUnits(subject.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	subjectsIndxs := binding.NewIntList()
	subjectsIndxs.Set(getIndexesFromList(subjects))
	unitsIndxs := binding.NewIntList()
	unitsIndxs.Set(getIndexesFromList(units))

	// Channels
	subjectChan := make(chan models.Subject)
	unitChan := make(chan models.Unit)

	// Build toolbar
	toolbar := widget.NewToolbar(
		// New subject
		widget.NewToolbarAction(theme.FolderNewIcon(), func() {
			newsub.Open(myApp, subjectChan)
			newSub := <-subjectChan // Wait for a sub
			// Append if not empty
			if newSub != (models.Subject{}) {
				subjects = append(subjects, newSub)
				subjectsIndxs.Set(getIndexesFromList(subjects))
			}
		}),
		// New Unit
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			newunit.Open(myApp, subject, unitChan)
			newUnit := <-unitChan // Wait for a unit
			// Append if not empty
			if newUnit != (models.Unit{}) {
				units = append(units, newUnit)
				unitsIndxs.Set(getIndexesFromList(units))
			}
		}),
		widget.NewToolbarSpacer(),
		// Settings
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settings.Open(myApp)
		}),
	)

	// Build list of subjects
	listSubjects := widget.NewListWithData(subjectsIndxs,
		func() fyne.CanvasObject {
			return widget.NewButton("subject", func() {})
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			index, _ := i.(binding.Int).Get()

			o.(*widget.Button).SetText(subjects[index].Name)
			o.(*widget.Button).OnTapped = func() {
				subject = subjects[index]
				units, err = db.GetUnits(subject.ID)
				if err != nil {
					fmt.Println("GetUnits:", err)
					return
				}
				unitsIndxs.Set(getIndexesFromList(units))
			}
		},
	)

	// Build list of units
	listUnits := widget.NewListWithData(unitsIndxs,
		func() fyne.CanvasObject {
			return widget.NewButton("unit", func() {})
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			index, _ := i.(binding.Int).Get()

			// Button name
			text := ""
			if units[index].Name != "" {
				text = fmt.Sprintf("%d: %s", units[index].Num, units[index].Name)
			} else {
				text = fmt.Sprintf("Unit %d", units[index].Num)
			}

			o.(*widget.Button).SetText(text)
			o.(*widget.Button).OnTapped = func() {
				mainWindow.Hide()
				helpers.LaunchEditor(myApp, subject, units[index])
				mainWindow.Show()
			}
		},
	)
	content := container.NewBorder(toolbar, nil, nil, nil, container.NewHSplit(listSubjects, listUnits))
	mainWindow.SetContent(content)
	mainWindow.Show()
}

func Show() {
	mainWindow.Show()
}
