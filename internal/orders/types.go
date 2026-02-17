package orders

import (
	"github.com/jackc/pgx/v5/pgtype"
	repo "github.com/majesticdrag0n/ecom/internal/adapters/postgresql/sqlc"
)

// PlaceOrderRequest represents the incoming request to create a new order along with its items.
type PlaceOrderRequest struct {
	CustomerName string           `json:"customer_name"`
	TotalAmount  pgtype.Numeric   `json:"total_amount"`
	Items        []OrderItemInput `json:"items"`
}

// OrderItemInput represents a single item in a place order request.
type OrderItemInput struct {
	ProductID pgtype.UUID    `json:"product_id"`
	Quantity  int32          `json:"quantity"`
	UnitPrice pgtype.Numeric `json:"unit_price"`
}

// PlaceOrderResponse wraps the created order and its items in a single response.
type PlaceOrderResponse struct {
	Order repo.Order       `json:"order"`
	Items []repo.OrderItem `json:"items"`
}
