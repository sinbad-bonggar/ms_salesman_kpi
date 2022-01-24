package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest/handler"
)

// HealthRouter a completely separate router for health check routes
func HealthRouter(h handler.IHealthHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/ping", h.Ping)
	r.Get("/health", h.HealthCheck)

	return r
}
