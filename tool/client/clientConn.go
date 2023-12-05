package client

import "time"

type ConnectionInfo struct {
	Timeout time.Duration
	SlaveId uint64
	Address string
}
