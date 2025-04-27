package main

import (
	"sync"
	"time"
)

type accessTracker struct {
	firstAccess    time.Time
	requestCounter int
}

type rateLimiter struct {
	requestLimit int
	duration     time.Duration
	trackers     map[string]*accessTracker
	mu           sync.RWMutex
}

func newRateLimiter(limit int, duration time.Duration) *rateLimiter {
	newAccessTracker := make(map[string])
	return &rateLimiter{
		requestLimit: limit,
		duration:     duration,
	}
}

func (rl *rateLimiter) checkTimeWindow(ip string) bool{
	rl.mu.RLock()
	defer rl.mu.RUnlock()
	timeDelta := time.Now().Sub(rl.trackers[ip].firstAccess)
	if timeDelta < rl.duration {
		return true
	}
	return false
}

func (rl *rateLimiter) resetTimeWindow(ip string) bool {
	time.Sleep(2 * time.Minute)
	rl.trackers[ip].firstAccess = 0
}

func (rl *rateLimiter) trackRequest (ip string) {
	rl.mu.Lock()
	existingIp := rl.trackers[ip]
	defer rl.mu.Unlock()

	if existingIp.requestCounter == 0 {
		rl.trackers[ip].firstAccess = time.Now()
		rl.trackers[ip].requestCounter = 1
	} 
	if existingIp.requestCounter < rl.requestLimit && rl.checkTimeWindow(ip) {
		rl.trackers[ip].requestCounter++
	} else {
		rl.resetTimeWindow(ip)
	}

}