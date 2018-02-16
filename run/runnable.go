package run

import (
	"github.com/mono83/xray"
)

// Runnable represents entities (functions), able to start
type Runnable func(xray.Ray) error
