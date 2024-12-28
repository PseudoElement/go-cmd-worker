package cmd

type WindowsCmdManager struct{}

func (this *WindowsCmdManager) CreateDir(dirName string) error {
	return nil
}

func (this *WindowsCmdManager) DeleteDir() error {
	return nil
}

func (this *WindowsCmdManager) GetCurrentDir() string {
	return ""
}

var _ OsCmdManager = (*WindowsCmdManager)(nil)
