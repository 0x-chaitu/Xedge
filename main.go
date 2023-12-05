package main

import (
	"fmt"
	"time"
	devicefactory "xedge/container/deviceFactory"
	edgecontainer "xedge/container/edgeContainer"
	"xedge/internal/poll"
	"xedge/tool/client"
	"xedge/tool/client/driver"
)

func main() {

	container := edgecontainer.EdgeContainer{FactoryMap: make(map[string]interface{})}

	config := client.ConnectionInfo{
		Timeout: 3 * time.Second,
		SlaveId: 1,
		Address: "127.0.0.1:502",
	}

	md, err := devicefactory.GetDeviceFactoryBuilder("modbus").Build(&container)
	if err != nil {
		fmt.Println(err, md)
	}
	modbusDevice := md.(driver.Driver)
	modbusDevice.SetConnectionInfo(&config)
	err = modbusDevice.Connect()
	if err != nil {
		fmt.Println(err)
	}
	subscription := poll.Subscribe(poll.ClientPoller{Driver: modbusDevice})
	for {
		go func() {
			for feed := range subscription.Updates() {
				fmt.Println(feed.Data)
			}
		}()
		time.Sleep(2 * time.Second)
	}

}
