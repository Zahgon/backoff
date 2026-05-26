package backoff

import (
	"context"
	"time"
)

// DefaultMaxElapsedTime sets a default limit for the total retry duration.
const DefaultMaxElapsedTime = 15 * time.Minute

// Operation is a function that attempts an operation and may be retried.
type Operation[T any] func() (T, error)

// Notify is a function called on operation error with the error and backoff duration.
type Notify func(error, time.Duration)

// retryOptions holds configuration settings for the retry mechanism.
type retryOptions struct {
	BackOff        BackOff       // Strategy for calculating backoff periods.
	Timer          timer         // Timer to manage retry delays.
	Notify         Notify        // Optional function to notify on each retry error.
	MaxTries       uint          // Maximum number of retry attempts.
	MaxElapsedTime time.Duration // Maximum total time for all retries.
}

type RetryOption func(*retryOptions)

// WithBackOff configures a custom backoff strategy.
func WithBackOff(b BackOff) RetryOption { _ = "STUB: not implemented"; return *new(RetryOption) }

// withTimer sets a custom timer for managing delays between retries.
func withTimer(t timer) RetryOption { _ = "STUB: not implemented"; return *new(RetryOption) }

// WithNotify sets a notification function to handle retry errors.
func WithNotify(n Notify) RetryOption { _ = "STUB: not implemented"; return *new(RetryOption) }

// WithMaxTries limits the number of all attempts.
func WithMaxTries(n uint) RetryOption { _ = "STUB: not implemented"; return *new(RetryOption) }

// WithMaxElapsedTime limits the total duration for retry attempts.
func WithMaxElapsedTime(d time.Duration) RetryOption {
	_ = "STUB: not implemented"
	return *new(RetryOption)
}

// Retry attempts the operation until success, a permanent error, or backoff completion.
// It ensures the operation is executed at least once.
//
// Returns the operation result or error if retries are exhausted or context is cancelled.
func Retry[T any](ctx context.Context, operation Operation[T], opts ...RetryOption) (T, error) {
	_ = "STUB: not implemented"
	// Initialize default retry options.
	return *new(T), nil
}

// Apply user-provided options to the default settings.

// Execute the operation.

// Handle permanent errors without retrying.

// Stop retrying if maximum tries exceeded.

// Stop retrying if context is cancelled.

// Calculate next backoff duration.

// Reset backoff if RetryAfterError is encountered.

// Stop retrying if maximum elapsed time exceeded.

// Notify on error if a notifier function is provided.

// Wait for the next backoff period or context cancellation.
