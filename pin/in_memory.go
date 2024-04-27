package pin

type InMemoryPINService struct {
	pin string
}

func NewInMemoryPINService() PINService {
	s := &InMemoryPINService{}
	s.Generate()
	return s
}

func (s *InMemoryPINService) CurrentPIN() string {
	return s.pin
}

func (s *InMemoryPINService) IsCorrect(pin string) bool {
	return pin == s.pin
}

func (s *InMemoryPINService) Generate() string {
	s.pin = generateRandomPIN()
	return s.pin
}
