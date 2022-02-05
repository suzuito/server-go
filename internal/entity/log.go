package entity

import "time"

type LogEntry struct {
	RemoteAddr string `json:"remoteAddr"`
	UserAgent  string `json:"userAgent"`

	URI         string    `json:"uri"`
	Method      string    `json:"method"`
	StartedAt   time.Time `json:"startedAt"`
	ResponsedAt time.Time `json:"responsedAt"`

	TargetURI         string    `json:"targetUri"`
	TargetMethod      string    `json:"targetMethod"`
	TargetStartedAt   time.Time `json:"targetStartedAt"`
	TargetResponsedAt time.Time `json:"targetResponsedAt"`
	TargetStatusCode  int       `json:"targetStatusCode"`
}
