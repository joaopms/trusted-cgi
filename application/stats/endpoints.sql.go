// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: endpoints.sql

package stats

import (
	"context"
	"time"
)

const addEndpointStat = `-- name: AddEndpointStat :exec
INSERT INTO endpoint_stat
(project, method, path, request_url, started_at, finished_at,
 request_size, request_headers, request_body,
 response_size, response_headers, response_body, status)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING id
`

type AddEndpointStatParams struct {
	Project         string    `json:"project"`
	Method          string    `json:"method"`
	Path            string    `json:"path"`
	RequestUrl      string    `json:"request_url"`
	StartedAt       time.Time `json:"started_at"`
	FinishedAt      time.Time `json:"finished_at"`
	RequestSize     int64     `json:"request_size"`
	RequestHeaders  string    `json:"request_headers"`
	RequestBody     []byte    `json:"request_body"`
	ResponseSize    int64     `json:"response_size"`
	ResponseHeaders string    `json:"response_headers"`
	ResponseBody    []byte    `json:"response_body"`
	Status          int64     `json:"status"`
}

func (q *Queries) AddEndpointStat(ctx context.Context, arg AddEndpointStatParams) error {
	_, err := q.db.ExecContext(ctx, addEndpointStat,
		arg.Project,
		arg.Method,
		arg.Path,
		arg.RequestUrl,
		arg.StartedAt,
		arg.FinishedAt,
		arg.RequestSize,
		arg.RequestHeaders,
		arg.RequestBody,
		arg.ResponseSize,
		arg.ResponseHeaders,
		arg.ResponseBody,
		arg.Status,
	)
	return err
}

const gCEndpointStats = `-- name: GCEndpointStats :exec
DELETE
FROM endpoint_stat
WHERE started_at < ?
`

func (q *Queries) GCEndpointStats(ctx context.Context, startedAt time.Time) error {
	_, err := q.db.ExecContext(ctx, gCEndpointStats, startedAt)
	return err
}

const getEndpointStat = `-- name: GetEndpointStat :one
SELECT id, project, method, path, request_url, started_at, finished_at, request_headers, request_body, request_size, response_headers, response_body, response_size, status
FROM endpoint_stat
WHERE id = ?
`

func (q *Queries) GetEndpointStat(ctx context.Context, id int64) (EndpointStat, error) {
	row := q.db.QueryRowContext(ctx, getEndpointStat, id)
	var i EndpointStat
	err := row.Scan(
		&i.ID,
		&i.Project,
		&i.Method,
		&i.Path,
		&i.RequestUrl,
		&i.StartedAt,
		&i.FinishedAt,
		&i.RequestHeaders,
		&i.RequestBody,
		&i.RequestSize,
		&i.ResponseHeaders,
		&i.ResponseBody,
		&i.ResponseSize,
		&i.Status,
	)
	return i, err
}

const listEndpointStats = `-- name: ListEndpointStats :many
SELECT id, project, method, path, request_url, started_at, finished_at, request_headers, request_body, request_size, response_headers, response_body, response_size, status
FROM endpoint_stat
ORDER BY id DESC
LIMIT ? OFFSET ?
`

type ListEndpointStatsParams struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

func (q *Queries) ListEndpointStats(ctx context.Context, arg ListEndpointStatsParams) ([]EndpointStat, error) {
	rows, err := q.db.QueryContext(ctx, listEndpointStats, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EndpointStat{}
	for rows.Next() {
		var i EndpointStat
		if err := rows.Scan(
			&i.ID,
			&i.Project,
			&i.Method,
			&i.Path,
			&i.RequestUrl,
			&i.StartedAt,
			&i.FinishedAt,
			&i.RequestHeaders,
			&i.RequestBody,
			&i.RequestSize,
			&i.ResponseHeaders,
			&i.ResponseBody,
			&i.ResponseSize,
			&i.Status,
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

const listProjectEndpointStats = `-- name: ListProjectEndpointStats :many
SELECT id, project, method, path, request_url, started_at, finished_at, request_headers, request_body, request_size, response_headers, response_body, response_size, status
FROM endpoint_stat
WHERE project = ?
ORDER BY id DESC
LIMIT ? OFFSET ?
`

type ListProjectEndpointStatsParams struct {
	Project string `json:"project"`
	Limit   int64  `json:"limit"`
	Offset  int64  `json:"offset"`
}

func (q *Queries) ListProjectEndpointStats(ctx context.Context, arg ListProjectEndpointStatsParams) ([]EndpointStat, error) {
	rows, err := q.db.QueryContext(ctx, listProjectEndpointStats, arg.Project, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EndpointStat{}
	for rows.Next() {
		var i EndpointStat
		if err := rows.Scan(
			&i.ID,
			&i.Project,
			&i.Method,
			&i.Path,
			&i.RequestUrl,
			&i.StartedAt,
			&i.FinishedAt,
			&i.RequestHeaders,
			&i.RequestBody,
			&i.RequestSize,
			&i.ResponseHeaders,
			&i.ResponseBody,
			&i.ResponseSize,
			&i.Status,
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

const listProjectSpecificEndpointStats = `-- name: ListProjectSpecificEndpointStats :many
SELECT id, project, method, path, request_url, started_at, finished_at, request_headers, request_body, request_size, response_headers, response_body, response_size, status
FROM endpoint_stat
WHERE project = ?
  AND method = ?
  AND path = ?
ORDER BY id DESC
LIMIT ? OFFSET ?
`

type ListProjectSpecificEndpointStatsParams struct {
	Project string `json:"project"`
	Method  string `json:"method"`
	Path    string `json:"path"`
	Limit   int64  `json:"limit"`
	Offset  int64  `json:"offset"`
}

func (q *Queries) ListProjectSpecificEndpointStats(ctx context.Context, arg ListProjectSpecificEndpointStatsParams) ([]EndpointStat, error) {
	rows, err := q.db.QueryContext(ctx, listProjectSpecificEndpointStats,
		arg.Project,
		arg.Method,
		arg.Path,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []EndpointStat{}
	for rows.Next() {
		var i EndpointStat
		if err := rows.Scan(
			&i.ID,
			&i.Project,
			&i.Method,
			&i.Path,
			&i.RequestUrl,
			&i.StartedAt,
			&i.FinishedAt,
			&i.RequestHeaders,
			&i.RequestBody,
			&i.RequestSize,
			&i.ResponseHeaders,
			&i.ResponseBody,
			&i.ResponseSize,
			&i.Status,
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