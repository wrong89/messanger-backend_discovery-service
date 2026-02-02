package redis

import (
	"context"
	"strings"
	"wrong89/messanger-backend_discovery-service/internal/models"

	"github.com/redis/go-redis/v9"
)

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

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) WriteServer(ctx context.Context, server models.Server) error {
	return s.db.Set(ctx, server.Title, server.GetAddress(), 0).Err()
}

func (s *Storage) GetServerByTitle(ctx context.Context, title string) (models.Server, error) {
	var result models.Server

	addr, err := s.db.Get(ctx, title).Result()
	if err != nil {
		return models.Server{}, err
	}

	args := strings.Split(addr, ":")
	result.Title = title
	result.Host = args[0]
	result.Port = args[1]

	return result, nil
}
