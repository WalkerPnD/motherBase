package main

import (
	"fmt"
	"runtime"

	"mjv.projects/motherBase/lib"
)

func main() {
	command := "milk"
	if runtime.GOOS == "windows" {
		command += ".bat"
	}
	fmt.Println(command)

	appErr := lib.StartQuickServ()
	if appErr != nil {
		panic(appErr)
	}
}
