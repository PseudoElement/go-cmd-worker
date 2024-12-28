package cmd

import (
	"os/exec"
	"strings"
)

type LinuxCmdManager struct{}

func (this *LinuxCmdManager) CreateDir(dirName string) error {
	return nil
}

func (this *LinuxCmdManager) DeleteDir() error {
	pwd := this.GetCurrentDir()
	splitted := strings.Split(pwd, "/")
	parentDirName := splitted[len(splitted)-1]

	err := exec.Command("sudo", "rm", "-rf", parentDirName).Run()
	if err != nil {
		WriteError(err)
	}

	return err
}

func (this *LinuxCmdManager) GetCurrentDir() string {
	out, _ := exec.Command("pwd").Output()
	splitted := strings.Split(string(out), " ")
	pwd := splitted[2]
	return pwd
}

var _ OsCmdManager = (*MacCmdManager)(nil)
