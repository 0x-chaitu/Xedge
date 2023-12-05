package poll

import (
	"log"
	"xedge/tool/client/driver"
	"xedge/tool/client/modbus"
)

// A Fetcher fetches Feed & Fetch returns a non-nil error.
type Fetcher interface {
	Fetch() (feeds []Feed, err error)
}

type ClientPoller struct {
	Driver driver.Driver
}

func (cp ClientPoller) Fetch() (feeds []Feed, err error) {
	cmd := new(modbus.CommandInfo)
	cmd.DataBank = modbus.HOLDING_REGISTERS
	cmd.StartingAddress = 0
	cmd.TotalAddress = 1
	data, err := cp.Driver.ReadValues(cmd)
	if err != nil {
		log.Println(err)
	}
	slice := []Feed{{Data: data}}
	return slice, nil

}
