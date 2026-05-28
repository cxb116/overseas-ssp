package config

import "time"

// NormalizeDuration accepts both Go-style duration values already stored as
// nanoseconds and plain integer millisecond values from legacy JSON configs.
func NormalizeDuration(value time.Duration, fallback time.Duration) time.Duration {
	if value <= 0 {
		return fallback
	}
	if value < time.Millisecond {
		return value * time.Millisecond
	}
	return value
}
