package main

import (
	"log"
	"net/http"

	"go-routine/internal/external"
	"go-routine/internal/httpserver"
	"go-routine/internal/repo"
	"go-routine/internal/service"
)

func main() {
	userRepo := repo.NewUserRepo()
	cardClient := external.NewCardClient()
	accountClient := external.NewAccountClient()

	authSvc := service.NewAuthService(userRepo)
	dashboardSvc := service.NewDashboardService(cardClient, accountClient, userRepo)

	h := httpserver.NewHandler(authSvc, dashboardSvc)

	// POST /login
	http.HandleFunc("/login", h.LoginAndDashboard)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
