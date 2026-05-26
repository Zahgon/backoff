package backoff

import (
	"sync"
	"time"
)

// Ticker holds a channel that delivers `ticks' of a clock at times reported by a BackOff.
//
// Ticks will continue to arrive when the previous operation is still running,
// so operations that take a while to fail could run in quick succession.
type Ticker struct {
	C        <-chan time.Time
	c        chan time.Time
	b        BackOff
	timer    timer
	stop     chan struct{}
	stopOnce sync.Once
}

// NewTicker returns a new Ticker containing a channel that will send
// the time at times specified by the BackOff argument. Ticker is
// guaranteed to tick at least once.  The channel is closed when Stop
// method is called or BackOff stops. It is not safe to manipulate the
// provided backoff policy (notably calling NextBackOff or Reset)
// while the ticker is running.
func NewTicker(b BackOff) *Ticker { _ = "STUB: not implemented"; return nil }

// Stop turns off a ticker. After Stop, no more ticks will be sent.
func (t *Ticker) Stop() { _ = "STUB: not implemented"; return }

func (t *Ticker) run() { _ = "STUB: not implemented"; return }

// Ticker is guaranteed to tick at least once.

// Prevent future ticks from being sent to the channel.

func (t *Ticker) send(tick time.Time) <-chan time.Time { _ = "STUB: not implemented"; return nil }
