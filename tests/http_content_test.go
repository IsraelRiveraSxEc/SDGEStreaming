package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"SDGEStreaming/internal/httpapi"
	"SDGEStreaming/internal/repositories"
	"SDGEStreaming/internal/services"
)

func TestGetAudiovisualEndpoint(t *testing.T) {
	SetupTestDB(t)
	defer TeardownTestDB()

	// Repositories
	userRepo := repositories.NewUserRepo()
	contentRepo := repositories.NewContentRepo()
	subscriptionRepo := repositories.NewSubscriptionRepo()
	playbackHistoryRepo := repositories.NewPlaybackHistoryRepo()
	favoriteRepo := repositories.NewFavoriteRepo()

	// Services
	userService := services.NewUserService(userRepo, subscriptionRepo)
	contentService := services.NewContentService(contentRepo)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo, userRepo)
	playbackService := services.NewPlaybackService(playbackHistoryRepo, favoriteRepo, contentRepo)

	// Router (igual que producción)
	mux := http.NewServeMux()
	httpapi.RegisterHandlers(mux, userService, contentService, subscriptionService, playbackService)

	req := httptest.NewRequest(http.MethodGet, "/api/content/audiovisual", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status esperado 200, recibido %d", w.Code)
	}
}
