package pin

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisPINService struct {
	db  *redis.Client
	key string
}

func NewRedisPINService(addr string, key string) (PINService, error) {
	s := &RedisPINService{
		db: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
		key: key,
	}

	if s.db.Exists(context.TODO(), key).Val() == 0 {
		if _, err := s.Generate(); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (s *RedisPINService) CurrentPIN() (string, error) {
	return s.db.Get(context.TODO(), s.key).Result()
}

func (s *RedisPINService) IsCorrect(pin string) (bool, error) {
	currentPIN, err := s.CurrentPIN()
	if err != nil {
		return false, err
	}
	return pin == currentPIN, nil
}

func (s *RedisPINService) Generate() (string, error) {
	p := generateRandomPIN()
	err := s.db.Set(context.TODO(), s.key, p, 0).Err()
	if err != nil {
		return "", err
	}
	return p, err
}
