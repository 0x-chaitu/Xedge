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
	*poller
	cmd        *modbus.CommandInfo
	pollsEvery *time.Ticker
	stopped    chan bool
}

func (s *Service) NewSubcription(driver driver.Driver, cmd *modbus.CommandInfo, pollEvery time.Duration) bool {
	ticker := time.NewTicker(pollEvery)
	stopped := make(chan bool, 1)
	poller := newPoller(driver)
	newSub := &Subscription{poller: poller, cmd: cmd, pollsEvery: ticker, stopped: stopped}
	err := newSub.subscribe()
	if err != nil {
		return false
	}

	go newSub.aggTicksIn(s)

	s.putSubscription("", newSub)
	return false
}

func (s *Subscription) subscribe() error {
	err := s.connect()
	if err != nil {
		return err
	}
	return nil
}

// aggregates individaul subscription ticks into service's ticker channel
func (sub *Subscription) aggTicksIn(s *Service) {
	for {
		select {
		case <-sub.pollsEvery.C:
			s.ticker <- sub
		case <-sub.stopped:
			return
		}
	}
}

func (*Service) UnSubscribe(s *Subscription) bool {
	if err := s.disConnect(); err != nil {
		return false
	}
	s.pollsEvery.Stop()
	close(s.stopped)
	return true

}
