package customers

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/majesticdrag0n/ecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	CreateCustomer(ctx context.Context, arg repo.CreateCustomerParams) (repo.Customer, error)
	GetCustomer(ctx context.Context, id pgtype.UUID) (repo.Customer, error)
	ListCustomers(ctx context.Context) ([]repo.Customer, error)
	UpdateCustomer(ctx context.Context, arg repo.UpdateCustomerParams) (repo.Customer, error)
	DeleteCustomer(ctx context.Context, id pgtype.UUID) error
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) CreateCustomer(ctx context.Context, arg repo.CreateCustomerParams) (repo.Customer, error) {
	return s.repo.CreateCustomer(ctx, arg)
}

func (s *svc) GetCustomer(ctx context.Context, id pgtype.UUID) (repo.Customer, error) {
	return s.repo.GetCustomer(ctx, id)
}

func (s *svc) ListCustomers(ctx context.Context) ([]repo.Customer, error) {
	return s.repo.ListCustomers(ctx)
}

func (s *svc) UpdateCustomer(ctx context.Context, arg repo.UpdateCustomerParams) (repo.Customer, error) {
	return s.repo.UpdateCustomer(ctx, arg)
}

func (s *svc) DeleteCustomer(ctx context.Context, id pgtype.UUID) error {
	return s.repo.DeleteCustomer(ctx, id)
}
