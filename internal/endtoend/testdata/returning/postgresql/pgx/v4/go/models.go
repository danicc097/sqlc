// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package querytest

import (
	"database/sql"
)

type User struct {
	Name sql.NullString
	ID   int32
}
