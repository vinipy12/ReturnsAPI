package main

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

type rateLimiter struct {
	requestLimit int
	duration     time.Duration
	trackers     map[string]*accessTracker
	mu           sync.RWMutex
}

func newRateLimiter(limit int, duration time.Duration) *rateLimiter {
	return &rateLimiter{
		requestLimit: limit,
		duration:     duration,
		trackers:     make(map[string]*accessTracker),
	}
}

func (rl *rateLimiter) checkTimeWindow(ip string) bool {
	rl.mu.RLock()
	defer rl.mu.RUnlock()
	timeDelta := time.Now().Sub(rl.trackers[ip].firstAccess)
	return timeDelta < rl.duration
}

func (rl *rateLimiter) resetTimeWindow(ip string) {
	rl.mu.Lock()
	rl.trackers[ip].requestCounter = 1
	rl.trackers[ip].firstAccess = time.Now()
	rl.mu.Unlock()
}

func (rl *rateLimiter) trackRequest(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	ipTracker, check := rl.trackers[ip]
	if !check {
		firstAccess := time.Now()
		requestCounter := 1
		rl.trackers[ip] = newAccessTracker(firstAccess, requestCounter)
		return true
	}

	if ipTracker.requestCounter < rl.requestLimit && rl.checkTimeWindow(ip) {
		rl.trackers[ip].requestCounter++
		return true
	} else {
		rl.resetTimeWindow(ip)
		return false
	}
}
