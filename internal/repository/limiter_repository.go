package repository

import (
	"sync"
	"time"
)

type limiterRepository struct {
	userData    map[string]int // Tracks user request counts
	globalCount int            // Tracks total global requests
	mu          sync.Mutex     // Mutex for thread safety
}

func NewLimiterRepository() *limiterRepository {
	repo := &limiterRepository{
		userData:    make(map[string]int),
		globalCount: 0,
	}

	// Start a goroutine to reset limits every second
	go repo.startResetTicker()

	return repo
}

func (r *limiterRepository) IncrementUserCount(clientID string) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.userData[clientID]++
	return r.userData[clientID]
}

func (r *limiterRepository) ResetUserCount(clientID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.userData[clientID] = 0
}

func (r *limiterRepository) GetUserCount(clientID string) int {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.userData[clientID]
}

func (r *limiterRepository) IncrementGlobalCount() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.globalCount++
	return r.globalCount
}

func (r *limiterRepository) ResetGlobalCount() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.globalCount = 0
}

func (r *limiterRepository) GetGlobalCount() int {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.globalCount
}

func (r *limiterRepository) startResetTicker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		r.mu.Lock()
		r.globalCount = 0
		for clientID := range r.userData {
			r.userData[clientID] = 0
		}
		r.mu.Unlock()
	}
}
