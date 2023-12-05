package poll

import "time"

// poll implements the Subscription interface.
type poll struct {
	fetcher Fetcher         // fetches items
	updates chan Feed       // sends items to the user
	closing chan chan error // for Close
}

// Subscribe returns a new Subscription that uses fetcher to fetch Items.
func Subscribe(fetcher Fetcher) Subscription {
	p := &poll{
		fetcher: fetcher,
		updates: make(chan Feed),       // for Updates
		closing: make(chan chan error), // for Close
	}
	go p.loop()
	return p
}

// Close implements Subscription.
func (p *poll) Close() error {
	errc := make(chan error)
	p.closing <- errc
	return <-errc
}

// Updates implements Subscription.
func (p *poll) Updates() <-chan Feed {
	return p.updates
}

func (p *poll) loop() {
	var err error
	var feeds []Feed
	ticker := time.NewTicker(5 * time.Second)

	for {
		var first Feed
		var updates chan Feed // HLupdates
		if len(feeds) > 0 {
			first = feeds[0]
			updates = p.updates // enable send case // HLupdates
		}
		select {
		case <-ticker.C:
			var fetched []Feed
			fetched, err = p.fetcher.Fetch()
			if err != nil {
				break
			}
			feeds = append(feeds, fetched...)
		case errc := <-p.closing: // HLchan
			errc <- err      // HLchan
			close(p.updates) // tells receiver we're done
			return
		case updates <- first:
			feeds = feeds[1:]
		}
	}
}
