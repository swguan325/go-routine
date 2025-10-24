package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"go-routine/internal/model"
	"go-routine/internal/service"
)

type Handler struct {
	authSvc      service.AuthService
	dashboardSvc service.DashboardService
}

func NewHandler(authSvc service.AuthService, dashboardSvc service.DashboardService) *Handler {
	return &Handler{
		authSvc:      authSvc,
		dashboardSvc: dashboardSvc,
	}
}

func (h *Handler) LoginAndDashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Step 1: login auth
	token, err := h.authSvc.Login(ctx, req.UserID, req.Password)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Step 2: build dashboard concurrently
	dashboard, err := h.dashboardSvc.GetDashboard(ctx, req.UserID)
	if err != nil {
		http.Error(w, "failed to build dashboard", http.StatusBadGateway)
		return
	}

	// Step 3: response
	resp := model.LoginResponse{
		Token:     token,
		Dashboard: *dashboard,
	}

	writeJSON(ctx, w, http.StatusOK, resp)
}

func writeJSON(ctx context.Context, w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
