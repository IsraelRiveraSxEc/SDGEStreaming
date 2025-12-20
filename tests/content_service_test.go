package tests

import (
	"testing"

	"SDGEStreaming/internal/repositories"
	"SDGEStreaming/internal/services"
)

func TestContentServiceCreate(t *testing.T) {
	SetupTestDB(t)
	defer TeardownTestDB()

	repo := repositories.NewContentRepo()
	service := services.NewContentService(repo)

	err := service.CreateAudiovisual(
		"Service Test",
		"Movie",
		"Comedy",
		90,
		"G",
		"Test synopsis",
		2024,
		"Director X",
	)

	if err != nil {
		t.Fatalf("error en service: %v", err)
	}
}
