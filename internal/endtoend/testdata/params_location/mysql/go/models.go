// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package querytest

import (
	"database/sql"
)

type Order struct {
	ID     int32
	Price  string
	UserID int32
}

type User struct {
	ID        int32
	FirstName string
	LastName  sql.NullString
	Age       int32
	JobStatus string
}
