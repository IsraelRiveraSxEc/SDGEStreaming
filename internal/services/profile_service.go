// internal/services/profile_service.go
package services

import (
	"SDGEStreaming/internal/models"
	"SDGEStreaming/internal/repositories"
	"fmt"
)

type ProfileService struct {
	profileRepo repositories.ProfileRepo
}

func NewProfileService(profileRepo repositories.ProfileRepo) *ProfileService {
	return &ProfileService{profileRepo: profileRepo}
}

// clasificamos edad igual que en UserService
func classifyAgeProfile(age int) string {
	switch {
	case age < 13:
		return "Niño"
	case age < 18:
		return "Adolescente"
	default:
		return "Adulto"
	}
}

func (s *ProfileService) CreateProfile(userID, age int, name string) (*models.Profile, error) {
	if age < 3 || age > 120 {
		return nil, fmt.Errorf("edad de perfil inválida")
	}
	if name == "" {
		return nil, fmt.Errorf("el nombre del perfil no puede estar vacío")
	}

	p := &models.Profile{
		UserID:    userID,
		Name:      name,
		Age:       age,
		AgeRating: classifyAgeProfile(age),
	}

	if err := s.profileRepo.Create(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProfileService) GetProfiles(userID int) ([]models.Profile, error) {
	return s.profileRepo.FindByUserID(userID)
}

func (s *ProfileService) DeleteProfile(profileID int) error {
	return s.profileRepo.Delete(profileID)
}

func (s *ProfileService) CountProfiles(userID int) (int, error) {
	return s.profileRepo.CountByUserID(userID)
}
