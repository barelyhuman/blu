package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {

	app := &App{}

	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		if _, err := os.Stat("/usr/local/bin/blueutil"); os.IsNotExist(err) {
			app.setBlueutilStatus(false, err.Error())
		} else {
			app.setBlueutilStatus(true, "")
		}
	}

	app.Install()
	app.appInstance.Run()
}

// CmdExecute - Execute a command on the shell
func CmdExecute(cmd string, args ...string) (bool, error) {
	if runtime.GOOS == "windows" {
		log.Fatal("Can't Execute this on a windows machine")
	} else {
		_, err := exec.Command(cmd, args...).Output()

		if err != nil {
			fmt.Printf("Error: %s \n", err)
			return false, err
		}
		return true, nil
	}
	return false, nil
}
