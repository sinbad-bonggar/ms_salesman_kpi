package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest/handler"
)

// OrderRouter a completely separate router for health check routes
func OrderRouter(h handler.IOrderHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.List)
	r.Get("/{id}", h.Detail)
	r.Post("/", h.Create)
	r.Patch("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)

	return r
}
