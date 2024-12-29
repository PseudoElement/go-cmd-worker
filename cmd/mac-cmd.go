package cmd

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type MacCmdManager struct{}

func (this *MacCmdManager) CreateDir(dirName string) error {
	err := exec.Command("mkdir", dirName).Run()
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (this *MacCmdManager) DeleteDir() error {
	pwd := this.GetCurrentDir()
	splitted := strings.Split(pwd, "/")
	parentDirName := "../" + splitted[len(splitted)-1]

	log.Println("pwd", pwd)
	log.Println("splitted", splitted)
	log.Println("parentDirName", parentDirName)

	err := exec.Command("sudo", "rm", "-r", parentDirName).Run()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (this *MacCmdManager) GetCurrentDir() string {
	ex, _ := os.Executable()
	path := filepath.Dir(ex)
	return path
}

var _ OsCmdManager = (*MacCmdManager)(nil)
