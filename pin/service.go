package pin

type PINService interface {
	CurrentPIN() string
	IsCorrect(pin string) bool
	Generate() string
}
