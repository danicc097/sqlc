// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createMemories = `-- name: CreateMemories :many
INSERT INTO memories (vampire_id)
SELECT
    unnest($1::uuid[]) AS vampire_id
RETURNING
    id, vampire_id, created_at, updated_at
`

func (q *Queries) CreateMemories(ctx context.Context, vampireID []uuid.UUID) ([]Memory, error) {
	rows, err := q.db.QueryContext(ctx, createMemories, pq.Array(vampireID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Memory
	for rows.Next() {
		var i Memory
		if err := rows.Scan(
			&i.ID,
			&i.VampireID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
