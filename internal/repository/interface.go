package repository

import (
	"context"
	"wrong89/messanger-backend_discovery-service/internal/models"
)

type Repo interface {
	WriteServer(ctx context.Context, server models.Server) error
	GetServerByTitle(ctx context.Context, title string) (models.Server, error)
}
