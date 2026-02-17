package products

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/majesticdrag0n/ecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repo.Product, error)
	GetProduct(ctx context.Context, id pgtype.UUID) (repo.Product, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error) {
	return s.repo.ListProducts(ctx)
}

func (s *svc) GetProduct(ctx context.Context, id pgtype.UUID) (repo.Product, error) {
	return s.repo.GetProduct(ctx, id)
}
