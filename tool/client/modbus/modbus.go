package modbus

import (
	"xedge/tool/client"

	"github.com/goburrow/modbus"
)

type ModBusClient struct {
	TCPClientHandler *modbus.TCPClientHandler

	RTUClientHandler *modbus.RTUClientHandler

	client modbus.Client
}

func (md *ModBusClient) Connect() error {
	error := md.TCPClientHandler.Connect()
	md.client = modbus.NewClient(md.TCPClientHandler)
	return error

}

func (md *ModBusClient) Disconnect() error {
	return nil
}

func (md *ModBusClient) ReadValues(info interface{}) ([]byte, error) {

	cmdInfo := info.(*CommandInfo)
	var result []byte
	var err error
	switch cmdInfo.DataBank {
	case COILS:
		result, err = md.client.ReadCoils(cmdInfo.StartingAddress, cmdInfo.TotalAddress)
	case HOLDING_REGISTERS:
		result, err = md.client.ReadHoldingRegisters(cmdInfo.StartingAddress, cmdInfo.TotalAddress)
	case INPUT_REGISTERS:
		result, err = md.client.ReadInputRegisters(cmdInfo.StartingAddress, cmdInfo.TotalAddress)
	case DISCRETE_INPUTS:
		result, err = md.client.ReadDiscreteInputs(cmdInfo.StartingAddress, cmdInfo.TotalAddress)
	}

	return result, err
}

func NewModBusClient() *ModBusClient {
	client := new(ModBusClient)
	return client
}

func (md *ModBusClient) SetConnectionInfo(info *client.ConnectionInfo) {

	handler := modbus.NewTCPClientHandler(info.Address)
	handler.SlaveId = byte(info.SlaveId)
	handler.Timeout = info.Timeout

	md.TCPClientHandler = handler
}
