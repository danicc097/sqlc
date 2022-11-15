// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/danicc097/sqlc/internal/sql/ast"
	"github.com/danicc097/sqlc/internal/sql/catalog"
)

var funcsPgStatStatements = []*catalog.Function{
	{
		Name: "pg_stat_statements",
		Args: []*catalog.Argument{
			{
				Name: "showtext",
				Type: &ast.TypeName{Name: "boolean"},
			},
		},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name:       "pg_stat_statements_info",
		Args:       []*catalog.Argument{},
		ReturnType: &ast.TypeName{Name: "record"},
	},
	{
		Name: "pg_stat_statements_reset",
		Args: []*catalog.Argument{
			{
				Name:       "userid",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "oid"},
			},
			{
				Name:       "dbid",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "oid"},
			},
			{
				Name:       "queryid",
				HasDefault: true,
				Type:       &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "void"},
	},
}

func PgStatStatements() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsPgStatStatements
	return s
}
