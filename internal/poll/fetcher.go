package poll

import (
	"log"
)

// A Fetcher fetches Feed & Fetch returns a non-nil error.
type Fetcher interface {
	Fetch() (feeds []Feed, err error)
}

func (cp *Subscription) Fetch() (feeds []Feed, err error) {
	data, err := cp.client.ReadValues(cp.cmd)
	if err != nil {
		log.Println(err)
	}
	slice := []Feed{{Data: data}}
	return slice, nil
}
