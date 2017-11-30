package env

import (
	"os/user"
)

// SystemUser contains information about system user
var SystemUser ArgSystemUser

// ArgSystemUser is wrapper over os.user, compatible with xray
type ArgSystemUser struct {
	user.User
}

// Int returns integer representation of pid
func (a ArgSystemUser) String() string { return a.Username }

// Name returns argument key (username)
func (ArgSystemUser) Name() string { return "username" }

// Value returns string representation of argument value
func (a ArgSystemUser) Value() string { return a.Username }

// Scalar returns raw representation of argument value. It can be scalar value or slice of scalar values.
func (a ArgSystemUser) Scalar() interface{} { return a.Username }

func init() {
	if u, err := user.Current(); err != nil {
		SystemUser = ArgSystemUser{User: *u}
	} else {
		SystemUser = ArgSystemUser{User: user.User{}}
	}
}
