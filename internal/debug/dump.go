package debug

import (
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/danicc097/sqlc/internal/opts"
)

var (
	Active bool
	Traced bool
	Debug  opts.Debug
)

func init() {
	Active = os.Getenv("SQLCDEBUG") != ""
	if Active {
		Debug = opts.DebugFromEnv()
		Traced = Debug.Trace != ""
	}
}

func Dump(n ...interface{}) {
	if Active {
		spew.Dump(n)
	}
}
