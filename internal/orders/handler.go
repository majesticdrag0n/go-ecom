package orders

import (
	"encoding/json"
	"log"
	"net/http"

	repo "github.com/majesticdrag0n/ecom/internal/adapters/postgresql/sqlc"
	jsonHelper "github.com/majesticdrag0n/ecom/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var req PlaceOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.PlaceOrder(r.Context(), req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonHelper.Write(w, http.StatusCreated, resp)
}

func (h *handler) AddOrderItem(w http.ResponseWriter, r *http.Request) {
	var params repo.AddOrderItemParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	item, err := h.service.AddOrderItem(r.Context(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonHelper.Write(w, http.StatusCreated, item)
}
