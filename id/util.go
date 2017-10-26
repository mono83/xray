package id

import (
	"crypto/sha256"
	"os"
	"strconv"
)

// Hostname contains local machine name
var Hostname string

// PID contains currently running application PID
var PID int

var gen20Base64Prefix []byte

func init() {
	var err error
	Hostname, err = os.Hostname()
	if err != nil {
		Hostname = "unknown"
	}
	PID = os.Getpid()

	// Generating host hash
	h := sha256.New()
	h.Write([]byte(Hostname + strconv.Itoa(PID)))
	gen20Base64Prefix = h.Sum(nil)[0:5]
}
