// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/danicc097/sqlc/internal/sql/ast"
	"github.com/danicc097/sqlc/internal/sql/catalog"
)

var funcsUnaccent = []*catalog.Function{
	{
		Name: "unaccent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "regdictionary"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "unaccent",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
}

func Unaccent() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsUnaccent
	return s
}
