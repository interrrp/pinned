package pin

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisPINService struct {
	db  *redis.Client
	key string
}

func NewRedisPINService(addr string, key string) PINService {
	s := &RedisPINService{
		db: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
		key: key,
	}

	if s.db.Exists(context.TODO(), key).Val() == 0 {
		s.Generate()
	}

	return s
}

func (s *RedisPINService) CurrentPIN() string {
	return s.db.Get(context.TODO(), s.key).Val()
}

func (s *RedisPINService) IsCorrect(pin string) bool {
	return pin == s.CurrentPIN()
}

func (s *RedisPINService) Generate() string {
	p := generateRandomPIN()
	s.db.Set(context.TODO(), s.key, p, 0)
	return p
}
