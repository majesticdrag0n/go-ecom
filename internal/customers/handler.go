package customers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func (h *handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var params repo.CreateCustomerParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	customer, err := h.service.CreateCustomer(r.Context(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonHelper.Write(w, http.StatusCreated, customer)
}

func (h *handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	parsed, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid customer ID", http.StatusBadRequest)
		return
	}

	id := pgtype.UUID{Bytes: parsed, Valid: true}

	customer, err := h.service.GetCustomer(r.Context(), id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonHelper.Write(w, http.StatusOK, customer)
}

func (h *handler) ListCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.service.ListCustomers(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonHelper.Write(w, http.StatusOK, customers)
}

func (h *handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	parsed, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid customer ID", http.StatusBadRequest)
		return
	}

	var params repo.UpdateCustomerParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	params.ID = pgtype.UUID{Bytes: parsed, Valid: true}

	customer, err := h.service.UpdateCustomer(r.Context(), params)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonHelper.Write(w, http.StatusOK, customer)
}

func (h *handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	parsed, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid customer ID", http.StatusBadRequest)
		return
	}

	id := pgtype.UUID{Bytes: parsed, Valid: true}

	if err := h.service.DeleteCustomer(r.Context(), id); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
