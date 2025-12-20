package main

import (
	"SDGEStreaming/internal/db"
	"SDGEStreaming/internal/httpapi"
	"SDGEStreaming/internal/models"
	"SDGEStreaming/internal/repositories"
	"SDGEStreaming/internal/security"
	"SDGEStreaming/internal/services"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := db.InitDB("sdgestreaming.db", "internal/db/migrations"); err != nil {
		log.Fatalf("Error fatal al iniciar la base de datos: %v", err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepo()
	contentRepo := repositories.NewContentRepo()
	subscriptionRepo := repositories.NewSubscriptionRepo()
	playbackHistoryRepo := repositories.NewPlaybackHistoryRepo()
	favoriteRepo := repositories.NewFavoriteRepo()

	userService := services.NewUserService(userRepo, subscriptionRepo)
	contentService := services.NewContentService(contentRepo)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo, userRepo)
	playbackService := services.NewPlaybackService(playbackHistoryRepo, favoriteRepo, contentRepo)

	createDefaultAdmin(userRepo)

	mux := http.NewServeMux()
	httpapi.RegisterHandlers(mux, userService, contentService, subscriptionService, playbackService)

	addr := ":8080"
	log.Printf("Servidor HTTP escuchando en http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Error al iniciar el servidor HTTP: %v", err)
	}
}

func createDefaultAdmin(userRepo repositories.UserRepo) {
	const adminEmail = "admin@sdge.com"
	const adminPassword = "admin123"

	existing, err := userRepo.FindByEmail(adminEmail)
	if err != nil {
		log.Printf("Error buscando admin por email: %v", err)
		return
	}
	if existing != nil {
		return
	}

	hashed, err := security.HashPassword(adminPassword)
	if err != nil {
		log.Printf("Error generando hash para admin: %v", err)
		return
	}

	now := time.Now()
	admin := &models.User{
		Name:         "Admin",
		Email:        adminEmail,
		Age:          30,
		PlanID:       3,
		AgeRating:    "Adulto",
		IsAdmin:      true,
		PasswordHash: hashed,
		CreatedAt:    now,
		LastLogin:    now,
	}

	if err := userRepo.Create(admin); err != nil {
		log.Printf("Error creando usuario admin por defecto: %v", err)
		return
	}

	log.Println("Usuario admin creado (admin@sdge.com / admin123).")
}
