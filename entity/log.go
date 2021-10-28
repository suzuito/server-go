package entity

import "time"

type LogEntry struct {
	RemoteAddr string
	UserAgent  string

	URI         string
	Method      string
	StartedAt   time.Time
	ResponsedAt time.Time
	StatusCode  int

	TargetURI         string
	TargetMethod      string
	TargetStartedAt   time.Time
	TargetResponsedAt time.Time
	TargetStatusCode  int
}
