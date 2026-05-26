package backoff

import "time"

type timer interface {
	Start(duration time.Duration)
	Stop()
	C() <-chan time.Time
}

// defaultTimer implements Timer interface using time.Timer
type defaultTimer struct {
	timer *time.Timer
}

// C returns the timers channel which receives the current time when the timer fires.
func (t *defaultTimer) C() <-chan time.Time {
	_ = "STUB: not implemented"

	// Start starts the timer to fire after the given duration
	return nil
}

func (t *defaultTimer) Start(duration time.Duration) { _ = "STUB: not implemented"; return }

// Stop is called when the timer is not used anymore and resources may be freed.
func (t *defaultTimer) Stop() { _ = "STUB: not implemented"; return }
