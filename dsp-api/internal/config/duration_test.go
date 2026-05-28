package config

import (
	"testing"
	"time"
)

func TestNormalizeDuration(t *testing.T) {
	tests := []struct {
		name     string
		value    time.Duration
		fallback time.Duration
		want     time.Duration
	}{
		{
			name:     "use fallback when unset",
			value:    0,
			fallback: time.Second,
			want:     time.Second,
		},
		{
			name:     "keep nanosecond encoded duration",
			value:    1300 * time.Millisecond,
			fallback: time.Second,
			want:     1300 * time.Millisecond,
		},
		{
			name:     "upgrade legacy millisecond integer",
			value:    1300,
			fallback: time.Second,
			want:     1300 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeDuration(tt.value, tt.fallback)
			if got != tt.want {
				t.Fatalf("NormalizeDuration(%v, %v) = %v, want %v", tt.value, tt.fallback, got, tt.want)
			}
		})
	}
}
