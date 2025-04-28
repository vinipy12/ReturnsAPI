package ratelimiter

import (
	"sync"
	"time"
)

type accessTracker struct {
	firstAccess    time.Time
	requestCounter int
}

func newAccessTracker(firstAccess time.Time, requestCounter int) *accessTracker {
	return &accessTracker{
		firstAccess:    firstAccess,
		requestCounter: requestCounter,
	}
}

// Trackers map grows until server restart, needs a cleanup for production environment

type RateLimiter struct {
	requestLimit int
	duration     time.Duration
	trackers     map[string]*accessTracker
	mu           sync.RWMutex
}

// Initializes a rate limiter for IP-based request tracking

func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		requestLimit: limit,
		duration:     duration,
		trackers:     make(map[string]*accessTracker),
	}
}

var InMemoryRateLimiter = NewRateLimiter(10, 1*time.Minute)

func (rl *RateLimiter) AllowRequest(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	_, check := rl.trackers[ip]
	if !check {
		firstAccess := time.Now()
		requestCounter := 1
		rl.trackers[ip] = newAccessTracker(firstAccess, requestCounter)
		return true
	}

	if time.Since(rl.trackers[ip].firstAccess) < rl.duration {
		if rl.trackers[ip].requestCounter < rl.requestLimit {
			rl.trackers[ip].requestCounter++
			return true
		}
		// Even if requestCounter == rl.requestLimit, the request should be blocked
		return false
	} else {
		rl.trackers[ip].requestCounter = 1
		rl.trackers[ip].firstAccess = time.Now()
		return true
	}
}
