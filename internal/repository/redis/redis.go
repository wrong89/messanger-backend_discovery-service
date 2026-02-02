package redis

import "github.com/redis/go-redis/v9"

type Storage struct {
	db *redis.Client
}

func NewStorage(connURL string) Storage {
	opts, err := redis.ParseURL(connURL)
	if err != nil {
		panic(err)
	}

	return Storage{
		db: redis.NewClient(opts),
	}
}


