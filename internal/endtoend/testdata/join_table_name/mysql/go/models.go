// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package querytest

type Bar struct {
	ID int64
}

type Foo struct {
	ID  int64
	Bar int64
}
