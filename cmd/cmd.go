package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	mytypes "github.com/pseudoelement/go-cmd-worker/types"
)

func GetCmdManager(operationSystem mytypes.OperationSystem) OsCmdManager {
	switch operationSystem {
	case mytypes.MACOS:
		return &MacCmdManager{}
	case mytypes.LINUX:
		return &LinuxCmdManager{}
	case mytypes.WINDOWS:
		return &WindowsCmdManager{}
	default:
		fmt.Printf("Unknown os type %s.\n", operationSystem)
		return nil
	}
}

func WriteError(err error) {
	pwd, _ := os.Getwd()
	pathToErrorFile := fmt.Sprintf("%s/logs/cmd-errors.txt", pwd)
	if _, e := os.Stat(pathToErrorFile); errors.Is(e, os.ErrNotExist) {
		f, _ := os.Create(pathToErrorFile)
		defer f.Close()
	}

	file, openErr := os.OpenFile(pathToErrorFile, os.O_APPEND|os.O_WRONLY, 0644)
	if openErr != nil {
		log.Fatal(openErr)
	}
	log.Println(err.Error())
	defer file.Close()

	msg := fmt.Sprintf("%s\n", err.Error())
	_, e := file.Write([]byte(msg))
	if e != nil {
		log.Println("Error in WriteError - ", e.Error())
	}
}
