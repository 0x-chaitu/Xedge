package devicefactory

import (
	"xedge/container"
	"xedge/tool/client/modbus"
)

type modbusFactory struct{}

func (mf *modbusFactory) Build(c container.Container) (DeviceInterface, error) {

	modbusClient := modbus.NewModBusClient()

	return modbusClient, nil
}
