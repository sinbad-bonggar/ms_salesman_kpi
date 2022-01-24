package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest/handler"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest/middleware"
)

// SellerCenterRouter a completely separate router for sinbad seller center routes
func SellerCenterRouter() http.Handler {
	r := chi.NewRouter()

	// sinbad seller center header
	r.Use(middleware.CheckSSCWebHeader)

	// register more seller-center routes over here ...

	return r
}

// SinbadAppRouter a completely separate router for sinbad app routes
func SinbadAppRouter(oh handler.IOrderHandler) http.Handler {
	r := chi.NewRouter()

	// sinbad app center header
	r.Use(middleware.CheckSinbadAppHeader)

	// working day routes
	r.Mount("/orders", OrderRouter(oh))

	// register more sinbad-app routes over here ...

	return r
}

// AgentAppRouter a completely separate router for agent app routes
func AgentAppRouter() http.Handler {
	r := chi.NewRouter()

	// agent app header
	r.Use(middleware.CheckAgentAppHeader)

	// register more agent-app routes over here ...

	return r
}

// AdminPanelRouter a completely separate router for admin panel routes
func AdminPanelRouter() http.Handler {
	r := chi.NewRouter()

	// admin panel header
	r.Use(middleware.CheckAPWebHeader)

	// register more admin-panel routes over here ...

	return r
}
