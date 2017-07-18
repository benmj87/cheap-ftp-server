package ftpserver

import (
	"os"
	"time"
)

// FTPFileInfo implements github.com/goftp/server/FileInfo
type FTPFileInfo struct{
	fi os.FileInfo
}

// NewFTPFileInfo returns a new instance of an FTP File info
func NewFTPFileInfo(fileInfo os.FileInfo) *FTPFileInfo {
	return &FTPFileInfo{
		fi: fileInfo,
	}
}

// Name base name of the file
func (fi *FTPFileInfo) Name() string {
	return fi.fi.Name()
} 

// Size length in bytes for regular files; system-dependent for others
func (fi *FTPFileInfo) Size() int64 {
	return fi.fi.Size()
}

// Mode file mode bits
func (fi *FTPFileInfo) Mode() os.FileMode {
	return fi.fi.Mode()
}  

// ModTime modification time
func (fi *FTPFileInfo) ModTime() time.Time {
	return fi.fi.ModTime()
}

// IsDir abbreviation for Mode().IsDir()
func (fi *FTPFileInfo) IsDir() bool {
	return fi.fi.IsDir()
}

// Sys underlying data source (can return nil)
func (fi *FTPFileInfo) Sys() interface{} {
	return fi.fi.Sys()
}

// Owner returns owner permissions on file ownership
func (fi *FTPFileInfo) Owner() string {
	return "test"
}

// Group returns group permissions on file ownership
func (fi *FTPFileInfo) Group() string {
	return "test"
}