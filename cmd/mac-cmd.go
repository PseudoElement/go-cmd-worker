package cmd

import (
	"log"
	"os/exec"
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

	log.Println("splitted", splitted)
	log.Println("parentDirName", parentDirName)

	err := exec.Command("sudo", "rm", "-r", parentDirName).Run()
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (this *MacCmdManager) GetCurrentDir() string {
	pwd, _ := exec.Command("pwd").Output()
	return string(pwd)
}

var _ OsCmdManager = (*MacCmdManager)(nil)
