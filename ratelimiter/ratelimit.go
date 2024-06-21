package ratelimiter

import (
	"errors"
	"sync"
	"time"
)

type Limit struct {
	Count    int
	Duration time.Duration
}

type RateLimiter struct {
	limits  map[string]Limit
	actions map[string]map[string][]time.Time
	mu      sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limits:  make(map[string]Limit),
		actions: make(map[string]map[string][]time.Time),
	}
}

// SetLimit set limit for action
func (rl *RateLimiter) SetLimit(action string, count int, duration time.Duration) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.limits[action] = Limit{Count: count, Duration: duration}
	if _, exists := rl.actions[action]; !exists {
		rl.actions[action] = make(map[string][]time.Time)
	}
}

// Allow check if user exceeded the limit
func (rl *RateLimiter) Allow(action, key string) (bool, error) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limit, exists := rl.limits[action]
	if !exists {
		return false, errors.New("limit not set for action")
	}

	actions, exists := rl.actions[action]
	if !exists {
		return false, errors.New("action not set")
	}

	now := time.Now()
	timestamps := actions[key]
	var newTimestamps []time.Time

	for _, timestamp := range timestamps {
		if now.Sub(timestamp) < limit.Duration {
			newTimestamps = append(newTimestamps, timestamp)
		}
	}

	// delete old data
	actions[key] = newTimestamps

	if len(newTimestamps) >= limit.Count {
		return false, nil
	}

	newTimestamps = append(newTimestamps, now)
	actions[key] = newTimestamps
	return true, nil
}
