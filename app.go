package main

import (
	"fyne.io/fyne"
	fyneApp "fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// AppButtons - button that the player or app uses that are dynamic and need changing
type AppButtons struct {
	likeButton *widget.Button
}

// AppLabels - labels that the player or app uses that are dynamic and need changing
type AppLabels struct {
	blueUtilStatus  *widget.Label
	errorStatusText *widget.Label
}

// AppWindows - all windows that the app might create or need
type AppWindows struct {
	mainWindow    fyne.Window
	warningWindow fyne.Window
}

// App - Base app wrapper with needed pointers and refs to handle app state
type App struct {
	appInstance   fyne.App
	windows       AppWindows
	labels        AppLabels
	blueUtilState int
	buttons       AppButtons
}

// Install - initial setup for the App
func (app *App) Install() {
	app.appInstance = fyneApp.NewWithID("im.reaper.blu")
}

func (app *App) setBlueutilStatus(found bool, err string) {
	labelText := "Blueutil not found, install using `brew install blueutil`"
	app.blueUtilState = 0
	if found {
		app.blueUtilState = 1
		labelText = "Blueutil Found."
	}
	app.labels.blueUtilStatus = widget.NewLabel(labelText)
	app.labels.errorStatusText = widget.NewLabel(err)
}

// DrawMainWindow - draw the main window
func (app *App) DrawMainWindow() {
	app.windows.mainWindow = app.appInstance.NewWindow("Blu")
	app.windows.mainWindow.Resize(fyne.NewSize(300, 40))
	resetButton := widget.NewButton("Reset Bluetooth", func() {
		CmdExecute(blueUtilPath, "-p", "0")
		CmdExecute("sleep", "2")
		CmdExecute(blueUtilPath, "-p", "1")
	})
	if app.blueUtilState == 0 {
		resetButton.Disable()
	}
	app.windows.mainWindow.SetContent(
		widget.NewVBox(
			app.labels.blueUtilStatus,
			resetButton,
		),
	)
}

// ShowMainWindow - show the main window
func (app *App) ShowMainWindow() {
	app.windows.mainWindow.Show()
}

// DrawWarningWindow - Draw the warning window with the given title and message
func (app *App) DrawWarningWindow(title string, message string) {
	app.windows.warningWindow = app.appInstance.NewWindow(title)
	warningLabel := widget.NewLabel(message)
	app.windows.warningWindow.SetContent(
		widget.NewVBox(
			warningLabel,
		),
	)
}

// ShowWarningWindow - show the main window
func (app *App) ShowWarningWindow() {
	app.windows.warningWindow.Show()
}
