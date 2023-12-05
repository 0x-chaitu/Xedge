package edgecontainer

import devicefactory "xedge/container/deviceFactory"

type EdgeContainer struct {
	FactoryMap map[string]interface{}
}

func (ec *EdgeContainer) InitApp() error {
	return nil
}

func (ec *EdgeContainer) BuildUseCase(code string) (interface{}, error) {
	return devicefactory.GetDeviceFactoryBuilder(code).Build(ec)
}

func (ec *EdgeContainer) Get(code string) (interface{}, bool) {
	value, found := ec.FactoryMap[code]
	return value, found
}

func (ec *EdgeContainer) Put(code string, value interface{}) {
	ec.FactoryMap[code] = value
}
