package products

import (
	"log/slog"
	"net/http"

	"github.com/kopyn00/go-booking-api/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListProducts(r.Context())
	if err != nil {
		slog.Error("ProductList:", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := []string{"Tort czekoladowy", "Tort truskawkowy"}
	json.Write(w, http.StatusOK, products)
}
