package pin

type PINService interface {
	CurrentPIN() (string, error)
	IsCorrect(pin string) (bool, error)
	Generate() (string, error)
}
