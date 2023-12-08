package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	edgecontainer "xedge/container/edgeContainer"
	"xedge/internal/poll"
	"xedge/server/router"
	"xedge/tool/client"
	"xedge/tool/client/driver"
	"xedge/tool/client/modbus"
)

func main() {

	container := edgecontainer.EdgeContainer{FactoryMap: make(map[string]interface{})}
	subsriptionService, _ := poll.NewSubscriptionService()

	config, err := client.NewConnectionInfo(3*time.Second, 1, "192.168.0.105:502")
	if err != nil {
		log.Println(err)
	}
	cmd := new(modbus.CommandInfo)
	cmd.DataBank = modbus.HOLDING_REGISTERS
	cmd.StartingAddress = 0
	cmd.TotalAddress = 1

	md, err := container.BuildUseCase("modbus")
	if err != nil {
		fmt.Println(err, md)
	}
	modbusDevice := md.(driver.Driver)
	modbusDevice.SetConnectionInfo(config)
	err = modbusDevice.Connect()
	if err != nil {
		fmt.Println(err)
	}
	go subsriptionService.StartService()
	subsriptionService.NewSubcription(modbusDevice, cmd, 3*time.Second)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	defer close(ch)
	go func() {
		for {
			go func() {
				for feed := range subsriptionService.Updates() {
					fmt.Println(feed.Data)
				}
			}()
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		time.Sleep(5 * time.Second)
		// s, _ := subsriptionService.GetSubscription("")
		// if s != nil {
		// 	subsriptionService.UnSubscribe(s)
		// }
	}()

	r := router.NewRouter()
	r.RunServer()

	log.Println(1)
	<-ch
	subsriptionService.Close()

}
