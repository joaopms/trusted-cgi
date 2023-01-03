// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package stats

import (
	"time"
)

type CronStat struct {
	ID         int64     `json:"id"`
	Project    string    `json:"project"`
	Expression string    `json:"expression"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
	Error      string    `json:"error"`
}

type EndpointStat struct {
	ID              int64     `json:"id"`
	Project         string    `json:"project"`
	Method          string    `json:"method"`
	Path            string    `json:"path"`
	RequestUrl      string    `json:"request_url"`
	StartedAt       time.Time `json:"started_at"`
	FinishedAt      time.Time `json:"finished_at"`
	RequestHeaders  string    `json:"request_headers"`
	RequestBody     []byte    `json:"request_body"`
	RequestSize     int64     `json:"request_size"`
	ResponseHeaders string    `json:"response_headers"`
	ResponseBody    []byte    `json:"response_body"`
	ResponseSize    int64     `json:"response_size"`
	Status          int64     `json:"status"`
}

type LambdaStat struct {
	ID          int64     `json:"id"`
	Project     string    `json:"project"`
	Name        string    `json:"name"`
	StartedAt   time.Time `json:"started_at"`
	FinishedAt  time.Time `json:"finished_at"`
	Environment string    `json:"environment"`
	Error       string    `json:"error"`
}
