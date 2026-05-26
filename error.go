package backoff

import (
	"time"
)

// PermanentError signals that the operation should not be retried.
type PermanentError struct {
	Err error
}

// Permanent wraps the given err in a *PermanentError.
func Permanent(err error) error { _ = "STUB: not implemented"; return nil }

// Error returns a string representation of the Permanent error.
func (e *PermanentError) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap returns the wrapped error.
func (e *PermanentError) Unwrap() error {
	_ = "STUB: not implemented"

	// RetryAfterError signals that the operation should be retried after the given duration.
	return nil
}

type RetryAfterError struct {
	Duration time.Duration
}

// RetryAfter returns a RetryAfter error that specifies how long to wait before retrying.
func RetryAfter(seconds int) error { _ = "STUB: not implemented"; return nil }

// Error returns a string representation of the RetryAfter error.
func (e *RetryAfterError) Error() string { _ = "STUB: not implemented"; return "" }
