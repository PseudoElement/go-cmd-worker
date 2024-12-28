package cmd

type OsCmdManager interface {
	DeleteDir() error
	CreateDir(dirName string) error
	GetCurrentDir() string
}
