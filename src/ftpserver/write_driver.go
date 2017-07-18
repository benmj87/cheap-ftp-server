package ftpserver

import (
	"fmt"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"
	"github.com/goftp/server"
	"os"
)

// WriteDriver implements a WriteOnly ftp server
type WriteDriver struct {
	// root holds the root directory
	root string
}

// NewWriteDriver returns a new write driver
func NewWriteDriver(rootpath string) (*WriteDriver) {
	return &WriteDriver{
		root: rootpath,
	}
}

// Init init
func (wd *WriteDriver) Init(c *server.Conn) {
	fmt.Printf("Root dir: %v\n", wd.root)
	fmt.Printf("Init called for %v:%v\n", c.PublicIp(), c.PassivePort())
}

// Stat returns fileinfo for the given path
// params  - a file path
// returns - a time indicating when the requested path was last modified
//         - an error if the file doesn't exist or the user lacks
//           permissions
func (wd *WriteDriver) Stat(fp string) (server.FileInfo, error) {
	path := filepath.Join(wd.root, fp)
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return NewFTPFileInfo(fi), nil
}

// ChangeDir changes the current directory
// params  - path
// returns - true if the current user is permitted to change to the
//           requested path
func (wd *WriteDriver) ChangeDir(path string) error {
	wd.root = filepath.Join(wd.root, path)
	return nil
}

// ListDir lists the contents of the current directory and calls run
// params  - path, function on file or subdir found
// returns - error
//           path
func (wd *WriteDriver) ListDir(path string, run func(server.FileInfo) error) error {
	fmt.Printf("Root dir: %v\n", wd.root)
	files, err := ioutil.ReadDir(wd.root)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	for _,file := range files {
		fi := NewFTPFileInfo(file)

		err = run(fi)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return err
		}
	}

	return nil
}

// DeleteDir deletes the current directory
// params  - path
// returns - true if the directory was deleted
func (wd *WriteDriver) DeleteDir(path string) error {
	return errors.New("Not implemented")
}

// DeleteFile deletes the current file
// params  - path
// returns - true if the file was deleted
func (wd *WriteDriver) DeleteFile(path string) error {
	return errors.New("Not implemented")
}

// Rename renames the file/dir to the new dest
// params  - from_path, to_path
// returns - true if the file was renamed
func (wd *WriteDriver) Rename(path string, newpath string) error {
	return errors.New("Not implemented")
}

// MakeDir creates the directory
// params  - path
// returns - true if the new directory was created
func (wd *WriteDriver) MakeDir(name string) error {
	return errors.New("Not implemented")
}

// GetFile returns the length and a ReadCloser for the path
// params  - path
// returns - a string containing the file data to send to the client
func (wd *WriteDriver) GetFile(path string, size int64) (int64, io.ReadCloser, error) {
	return 0, nil, errors.New("Not implemented")
}

// PutFile writes a file to disk, returns the number of bytes written
// params  - destination path, an io.Reader containing the file data
// returns - true if the data was successfully persisted
func (wd *WriteDriver) PutFile(path string, reader io.Reader, b bool) (int64, error) {
	filename := filepath.Join(wd.root, path)
	fmt.Printf("Creating file %v\n", filename)
	
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return 0, err
	}

	written, err := io.Copy(file, reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return written, err
	}

	return written, nil
}