package ftpserver

import (
	"fmt"
)

// Auth implements github.com/goftp/server#Auth and always returns true
type Auth struct {
	// nop
}

// NewAuth returns a new default instance of Auth
func NewAuth() *Auth {
	return &Auth{}
}

// CheckPasswd always returns true
func (a *Auth) CheckPasswd(username string, password string) (bool, error) {
	fmt.Printf("Authenticating %v\n", username)
	return true, nil
}