package ratelimiter_test

import (
	"testing"
	"time"

	"github.com/Spargwy/ratelimiter/ratelimiter"
)

func TestRateLimiter(t *testing.T) {
	rl := ratelimiter.NewRateLimiter()
	rl.SetLimit("user_message", 5, time.Second)
	rl.SetLimit("ip_request", 100, time.Minute)

	userID := "testuser"
	ip := "127.0.0.1"

	for i := 0; i < 5; i++ {
		allowed, err := rl.Allow("user_message", userID)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !allowed {
			t.Errorf("Expected to allow user message, attempt %d", i+1)
		}
	}

	allowed, err := rl.Allow("user_message", userID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if allowed {
		t.Errorf("Expected to deny user message")
	}

	for i := 0; i < 100; i++ {
		allowed, err := rl.Allow("ip_request", ip)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !allowed {
			t.Errorf("Expected to allow IP request, attempt %d", i+1)
		}
	}

	allowed, err = rl.Allow("ip_request", ip)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if allowed {
		t.Errorf("Expected to deny IP request")
	}
}
