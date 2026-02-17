package orders

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	repo "github.com/majesticdrag0n/ecom/internal/adapters/postgresql/sqlc"
)

type Service interface {
	PlaceOrder(ctx context.Context, req PlaceOrderRequest) (PlaceOrderResponse, error)
	AddOrderItem(ctx context.Context, arg repo.AddOrderItemParams) (repo.OrderItem, error)
}

type svc struct {
	repo repo.Querier
	db   *pgx.Conn // raw connection needed to begin transactions
}

func NewService(repo repo.Querier, db *pgx.Conn) Service {
	return &svc{repo: repo, db: db}
}

// PlaceOrder creates an order and its items inside a single transaction.
// If any step fails, the entire transaction is rolled back.
func (s *svc) PlaceOrder(ctx context.Context, req PlaceOrderRequest) (PlaceOrderResponse, error) {
	// Begin a transaction
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return PlaceOrderResponse{}, fmt.Errorf("begin transaction: %w", err)
	}

	// Ensure rollback on any error — safe to call even after a successful commit
	defer tx.Rollback(ctx)

	// Use a transaction-scoped query instance
	qtx := repo.New(tx)

	// 1. Create the order
	order, err := qtx.PlaceOrder(ctx, repo.PlaceOrderParams{
		CustomerName: req.CustomerName,
		TotalAmount:  req.TotalAmount,
	})
	if err != nil {
		return PlaceOrderResponse{}, fmt.Errorf("place order: %w", err)
	}

	// 2. Add each order item and deduct product quantity
	var items []repo.OrderItem
	for _, itemInput := range req.Items {
		item, err := qtx.AddOrderItem(ctx, repo.AddOrderItemParams{
			OrderID:   order.ID,
			ProductID: itemInput.ProductID,
			Quantity:  itemInput.Quantity,
			UnitPrice: itemInput.UnitPrice,
		})
		if err != nil {
			return PlaceOrderResponse{}, fmt.Errorf("add order item: %w", err)
		}
		items = append(items, item)

		// Deduct product stock — fails if insufficient quantity
		_, err = qtx.UpdateProductQuantity(ctx, repo.UpdateProductQuantityParams{
			ID:       itemInput.ProductID,
			Quantity: itemInput.Quantity,
		})
		if err != nil {
			return PlaceOrderResponse{}, fmt.Errorf("insufficient stock for product %v: %w", itemInput.ProductID, err)
		}
	}

	// 3. Commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return PlaceOrderResponse{}, fmt.Errorf("commit transaction: %w", err)
	}

	return PlaceOrderResponse{
		Order: order,
		Items: items,
	}, nil
}

func (s *svc) AddOrderItem(ctx context.Context, arg repo.AddOrderItemParams) (repo.OrderItem, error) {
	return s.repo.AddOrderItem(ctx, arg)
}
