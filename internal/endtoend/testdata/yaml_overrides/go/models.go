// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package override

import (
	"github.com/danicc097/sqlc-testdata/pkg"
	"github.com/lib/pq"
)

type Foo struct {
	Other   string
	Total   int64
	Tags    []string
	ByteSeq []byte
	Retyped pkg.CustomType
	Langs   pq.StringArray
}
