package services

type PingService struct{}

// NewPingService creates a new PingService
func NewPingService() *PingService {
	return &PingService{}
}

func (u *PingService) Ping() (string, error) {
	return "pong", nil
}
