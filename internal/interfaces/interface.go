package interfaces

// LimiterRepository defines methods for tracking user and global limits.
type LimiterRepository interface {
	IncrementUserCount(clientID string) int
	ResetUserCount(clientID string)
	GetUserCount(clientID string) int

	IncrementGlobalCount() int
	ResetGlobalCount()
	GetGlobalCount() int
}

type LimiterService interface {
	HandleRequest(clientID string) (bool, bool) // Returns (userAllowed, globalAllowed)
}
