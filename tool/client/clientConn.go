package client

import "time"

type ConnectionInfo struct {
	Timeout time.Duration
	SlaveId uint64
	Address string
}

func NewConnectionInfo(timeout time.Duration, slaveId uint64, address string) (*ConnectionInfo, error) {
	conn := new(ConnectionInfo)
	conn.Timeout = timeout
	conn.SlaveId = slaveId
	conn.Address = address
	return conn, nil
}
