package id

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"math/rand"
	"sync"
	"time"
)

var generator20Mutex sync.Mutex
var generator20Inc int64
var generator20Rng = rand.New(rand.NewSource(time.Now().Unix()))

// Generator20Base64 always produces 20-bytes ray IDs but be aware of special symbols + and /
func Generator20Base64() string {
	b2 := make([]byte, 2)
	b8 := make([]byte, 8)
	buf := bytes.NewBuffer(nil)

	generator20Mutex.Lock()
	generator20Inc++
	inc := uint16(generator20Inc % 32768)
	rng := uint16(generator20Rng.Int() % 32768)
	generator20Mutex.Unlock()

	binary.BigEndian.PutUint16(b2, rng)
	buf.Write(b2)
	binary.BigEndian.PutUint64(b8, uint64(time.Now().UnixNano()/int64(time.Millisecond)))
	buf.Write(b8[2:])
	binary.BigEndian.PutUint16(b2, inc)
	buf.Write(b2)
	buf.Write(gen20Base64Prefix)

	return base64.StdEncoding.EncodeToString([]byte(buf.Bytes()))
}
