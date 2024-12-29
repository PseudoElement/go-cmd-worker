package main

import (
	"log"
	"runtime"

	"github.com/pseudoelement/go-cmd-worker/cmd"
	mytypes "github.com/pseudoelement/go-cmd-worker/types"
)

func detectUserOS() mytypes.OperationSystem {
	val := runtime.GOOS
	if val == "darwin" {
		return mytypes.MACOS
	}
	if val == "windows" {
		return mytypes.WINDOWS
	}
	if val == "linux" {
		return mytypes.LINUX
	}
	return val
}

func main() {
	userOS := detectUserOS()
	cmdManager := cmd.GetCmdManager(userOS)
	err := cmdManager.DeleteDir()
	if err != nil {
		log.Println(err)
	}
	er := cmdManager.CreateDir("my-dir")
	if er != nil {
		log.Println(er)
	}
}
