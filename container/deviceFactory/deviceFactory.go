package devicefactory

import (
	"xedge/container"
)

type DeviceInterface interface{}

var deviceFactoryMap = map[string]deviceFbInterface{
	"modbus": &modbusFactory{},
}

type deviceFbInterface interface {
	Build(c container.Container) (DeviceInterface, error)
}

func GetDeviceFactoryBuilder(key string) deviceFbInterface {
	return deviceFactoryMap[key]
}
