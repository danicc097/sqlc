package bundler

import (
	"runtime"

	"github.com/danicc097/sqlc/internal/info"
)

func projectMetadata() ([][2]string, error) {
	return [][2]string{
		{"sqlc_version", info.Version},
		{"go_version", runtime.Version()},
		{"goos", runtime.GOOS},
		{"goarch", runtime.GOARCH},
	}, nil
}
