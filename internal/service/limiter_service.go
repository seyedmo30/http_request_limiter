package service

import (
	"github.com/seyedmo30/http_request_limiter/internal/config"
	"github.com/seyedmo30/http_request_limiter/internal/interfaces"
)

type limiterService struct {
	repo   interfaces.LimiterRepository
	config config.Config
}

func NewLimiterService(repo interfaces.LimiterRepository, config config.Config) interfaces.LimiterService {
	return &limiterService{
		repo:   repo,
		config: config,
	}
}

func (s *limiterService) HandleRequest(clientID string) (bool, bool) {
	// Check and increment global count
	globalCount := s.repo.IncrementGlobalCount()
	if globalCount > s.config.GlobalRequestLimit {
		return true, false
	}

	// Check and increment user-specific count
	userCount := s.repo.IncrementUserCount(clientID)
	if userCount > s.config.UserRequestLimit {
		return false, true
	}

	return true, true
}
