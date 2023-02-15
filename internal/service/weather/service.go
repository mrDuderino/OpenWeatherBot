package weather

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Show() Weather {
	return weatherSum
}
