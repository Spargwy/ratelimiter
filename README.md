## Ratelimiter
- go get github.com/Spargwy/ratelimiter

example:
```
func main() {
    rl := ratelimiter.NewRateLimiter()
	rl.SetLimit("user_message", 5, time.Second)

	userID := "testuser"

	for i := 0; i < 5; i++ {
		allowed, err := rl.Allow("user_message", userID)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !allowed {
			t.Errorf("Expected to allow user message, attempt %d", i+1)
		}
	}
}
```