// genericdevice package represents low level client interfaces in order to have an unified way to
// handle connection, read data bytes value and write data bytes value

package driver

import "xedge/tool/client"

type Driver interface {
	Connect() error

	Disconnect() error

	ReadValues(info interface{}) ([]byte, error)

	SetConnectionInfo(info *client.ConnectionInfo)
}
