package cmd

import (
	"os/exec"
	"strings"
)

type MacCmdManager struct{}

func (this *MacCmdManager) CreateDir(dirName string) error {
	return nil
}

func (this *MacCmdManager) DeleteDir() error {
	pwd := this.GetCurrentDir()
	splitted := strings.Split(pwd, "/")
	parentDirName := splitted[len(splitted)-1]

	err := exec.Command("sudo", "rm", "-r", parentDirName).Run()
	if err != nil {
		WriteError(err)
	}

	return err
}

func (this *MacCmdManager) GetCurrentDir() string {
	pwd, _ := exec.Command("pwd").Output()
	return string(pwd)
}

var _ OsCmdManager = (*MacCmdManager)(nil)
