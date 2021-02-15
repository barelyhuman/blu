package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

const blueUtilPath = "/usr/local/bin/blueutil"

func main() {

	app := &App{}

	app.Install()

	if runtime.GOOS != "darwin" {
		// TODO: Show this as a popup on the main app to let the user know
		fmt.Println("Can't Execute this on a windows machine")
		app.DrawWarningWindow("Error:", "Can't Execute this on a windows machine")
		app.ShowWarningWindow()
	} else {
		app.Install()
		if _, err := os.Stat(blueUtilPath); os.IsNotExist(err) {
			app.setBlueutilStatus(false, err.Error())
		} else {
			app.setBlueutilStatus(true, "")
		}
		app.DrawMainWindow()
		app.ShowMainWindow()
	}

	app.appInstance.Run()
}

// CmdExecute - Execute a command on the shell
func CmdExecute(cmd string, args ...string) (bool, error) {
	if runtime.GOOS == "windows" {
		log.Fatal("Can't Execute this on a windows machine")
	}

	_, err := exec.Command(cmd, args...).Output()

	if err != nil {
		fmt.Printf("Error: %s \n", err)
		return false, err
	}
	return true, nil

}
