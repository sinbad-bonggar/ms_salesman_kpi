package handler

import (
	"net/http"

	"github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest/response"
)

// IHealthHandler ...
type IHealthHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct {
	response response.IResponseClient
}

// NewHealthHandler ...
func NewHealthHandler(r response.IResponseClient) IHealthHandler {
	return &healthHandler{
		response: r,
	}
}

// Ping checks that the API is running and is accessible
func (h *healthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	h.response.JSON(w, "Pong", nil, nil)
}

// HealthCheck endpoint verifies multiple items and responds with the status of the API and its dependencies
func (h *healthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// TODO: check all dependencies like db conn, and kafka.

	payload := response.HealthCheckMessage{
		ServiceName: "JENKINS_SERVICENAME",
		Version:     "JENKINS_GIT_TAG",
		CommitId:    "JENKINS_GIT_COMMIT",
		UpdatedAt:   "JENKINS_TIME",
		Status:      "pass",
	}

	h.response.JSON(w, "OK", payload, nil)
}
