package tests

import (
	"testing"

	"SDGEStreaming/internal/models"
	"SDGEStreaming/internal/repositories"
)

func TestFilterByAgeRating(t *testing.T) {
	SetupTestDB(t)
	defer TeardownTestDB()

	repo := repositories.NewContentRepo()

	err := repo.CreateAudiovisual(&models.AudiovisualContent{
		Title:     "Adult Movie",
		Type:      "Movie",
		Genre:     "Action",
		Duration:  100,
		AgeRating: "PG-13",
		IsAvailable: true,
		AverageRating: 0.0,		
	})
	if err != nil {
		t.Fatalf("error creando contenido: %v", err)
	}

	niño, _ := repo.FindAllAudiovisualAllowed("Niño")
	if len(niño) != 0 {
		t.Fatalf("Niño no debería ver contenido PG-13")
	}

	adolescente, _ := repo.FindAllAudiovisualAllowed("Adolescente")
	if len(adolescente) != 1 {
		t.Fatalf("Adolescente debería ver contenido PG-13")
	}

	adulto, _ := repo.FindAllAudiovisualAllowed("Adulto")
	if len(adulto) != 1 {
		t.Fatalf("Adulto debería ver todo el contenido")
	}
}
