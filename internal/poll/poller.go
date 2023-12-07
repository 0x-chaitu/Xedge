package poll

import (
	"xedge/tool/client/driver"
)

type poller struct {
	client    driver.Driver
	connected bool
}

func newPoller(d driver.Driver) *poller {
	return &poller{client: d, connected: false}
}

func (p *poller) connect() error {
	if err := p.client.Connect(); err != nil {
		p.connected = false
		return err
	}
	p.connected = true
	return nil
}

func (p *poller) disConnect() error {
	if p.connected {
		if err := p.client.Disconnect(); err != nil {
			return err
		}
		p.connected = false
	}
	return nil

}
