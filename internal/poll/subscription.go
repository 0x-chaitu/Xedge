package poll

// A Subscription delivers Feeds over a channel.  Close cancels the
// subscription, closes the Updates channel, and returns the last fetch error,
// if any.
type Subscription interface {
	Updates() <-chan Feed
	Close() error
}
