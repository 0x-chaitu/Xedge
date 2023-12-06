package poll

import (
	"time"
	"xedge/tool/client/driver"
	"xedge/tool/client/modbus"
)

// A Subscription delivers Feeds over a channel.  Close cancels the
// subscription, closes the Updates channel, and returns the last fetch error,
// if any.
type SubscriptionService interface {
	Updates() <-chan Feed
	Close() error
}

type Subscription struct {
	driver.Driver
	cmd *modbus.CommandInfo
	*time.Ticker
}

func (s *Service) NewSubcription(driver driver.Driver, cmd *modbus.CommandInfo, pollEvery time.Duration) {
	ticker := time.NewTicker(pollEvery)
	s.putSubscription("", &Subscription{Driver: driver, cmd: cmd, Ticker: ticker})
}
