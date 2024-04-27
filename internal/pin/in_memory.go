package pin

type InMemoryPINService struct {
	pin string
}

func NewInMemoryPINService() PINService {
	s := &InMemoryPINService{}
	s.Generate()
	return s
}

func (s *InMemoryPINService) CurrentPIN() (string, error) {
	return s.pin, nil
}

func (s *InMemoryPINService) IsCorrect(pin string) (bool, error) {
	return pin == s.pin, nil
}

func (s *InMemoryPINService) Generate() (string, error) {
	s.pin = generateRandomPIN()
	return s.pin, nil
}
