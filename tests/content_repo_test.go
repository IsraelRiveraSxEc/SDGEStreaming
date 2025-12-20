package tests

import (
	"testing"

	"SDGEStreaming/internal/models"
	"SDGEStreaming/internal/repositories"
)

func TestCreateAndListAudiovisual(t *testing.T) {
	SetupTestDB(t)
	defer TeardownTestDB()

	repo := repositories.NewContentRepo()

	content := &models.AudiovisualContent{
		Title:     "Test Movie",
		Type:      "Movie",
		Genre:     "Drama",
		Duration:  120,
		AgeRating: "PG-13",
		IsAvailable: true,
		AverageRating: 0,
	}

	err := repo.CreateAudiovisual(content)
	if err != nil {
		t.Fatalf("error creando contenido: %v", err)
	}

	list, err := repo.FindAllAudiovisual()
	if err != nil {
		t.Fatalf("error listando contenido: %v", err)
	}

	if len(list) != 1 {
		t.Fatalf("se esperaba 1 contenido, se obtuvo %d", len(list))
	}
}
