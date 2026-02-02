package redis

import "github.com/redis/go-redis/v9"

type Storage struct {
	db *redis.Client
}

func NewStorage(dbURL string) Storage {
	opts, err := redis.ParseURL(dbURL)
	if err != nil {
		panic(err)
	}

	return Storage{
		db: redis.NewClient(opts),
	}
}
