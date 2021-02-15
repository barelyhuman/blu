package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {

	app := &App{}

	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		executed := CmdExecute("which", "blueutil")
		if !executed {
			app.setBlueutilStatus(false)
		} else {
			app.setBlueutilStatus(true)
		}
	}

	app.Install()
	app.appInstance.Run()
}

// CmdExecute - Execute a command on the shell
func CmdExecute(cmd string, args ...string) bool {
	if runtime.GOOS == "windows" {
		log.Fatal("Can't Execute this on a windows machine")
	} else {
		_, err := exec.Command(cmd, args...).Output()

		if err != nil {
			fmt.Printf("Error: %s \n", err)
			return false
		}
		return true
	}
	return false
}
