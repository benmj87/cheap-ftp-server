package ftpserver

import (
	"github.com/goftp/server"
)

// DriverFactory provides an implementation of server.NewDriver returning the WriteDriver
type DriverFactory struct {
	// root is the root dir to pass through to the WriteDirver
	root string
}

// NewDriverFactory creates and returns a default instance of NewDriverFactory
func NewDriverFactory(rootpath string) (*DriverFactory) {
	return &DriverFactory{
		root: rootpath,
	}
}

// NewDriver creates a new WriteDriver
func (df *DriverFactory) NewDriver() (server.Driver, error) {
	driver := NewWriteDriver(df.root)
	return driver, nil
}