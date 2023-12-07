package poll

type Service struct {
	subscriptions map[string]*Subscription
	updates       chan Feed       // sends items to the user
	closing       chan chan error // for Close
	ticker        chan *Subscription
}

func NewSubscriptionService() (*Service, error) {
	subscriptions := make(map[string]*Subscription)
	ticker := make(chan *Subscription)
	updates := make(chan Feed)
	closing := make(chan chan error)
	return &Service{subscriptions: subscriptions, ticker: ticker, updates: updates, closing: closing}, nil
}

func (s *Service) putSubscription(key string, subscription *Subscription) {
	s.subscriptions[key] = subscription
}

func (s *Service) GetSubscription(key string) (*Subscription, error) {
	return s.subscriptions[key], nil
}

// Close implements Subscription.
func (s *Service) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}

// Updates implements Subscription.
func (s *Service) Updates() <-chan Feed {
	return s.updates
}

func (s *Service) StartService() {
	var err error
	var feeds []Feed

	for {
		var first Feed
		var updates chan Feed // HLupdates
		if len(feeds) > 0 {
			first = feeds[0]
			updates = s.updates // enable send case // HLupdates
		}
		select {
		case subscription := <-s.ticker:
			var fetched []Feed
			fetched, err = subscription.Fetch()
			if err != nil {
				break
			}
			feeds = append(feeds, fetched...)
		case errc := <-s.closing: // HLchan
			errc <- err      // HLchan
			close(s.updates) // tells receiver we're done
			return
		case updates <- first:
			feeds = feeds[1:]
		}
	}
}

func (s *Service) RestartAggregation() {
	for _, sub := range s.subscriptions {
		go func(sub *Subscription) {
			sub.aggTicksIn(s)
		}(sub)
	}

}
