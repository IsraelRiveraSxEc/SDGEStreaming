// internal/repositories/profile_repo.go
package repositories

import (
	"SDGEStreaming/internal/db"
	"SDGEStreaming/internal/models"
	"database/sql"
	"fmt"
)

type ProfileRepo interface {
	Create(profile *models.Profile) error
	FindByID(id int) (*models.Profile, error)
	FindByUserID(userID int) ([]models.Profile, error)
	CountByUserID(userID int) (int, error)
	Delete(id int) error
}

type sqliteProfileRepo struct{}

func NewProfileRepo() ProfileRepo {
	return &sqliteProfileRepo{}
}

func (r *sqliteProfileRepo) Create(profile *models.Profile) error {
	conn := db.GetDB()

	result, err := conn.Exec(`
		INSERT INTO profiles (user_id, name, age, age_rating)
		VALUES (?, ?, ?, ?)
	`, profile.UserID, profile.Name, profile.Age, profile.AgeRating)
	if err != nil {
		return fmt.Errorf("no se pudo crear el perfil: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	profile.ID = int(id)
	return nil
}

func (r *sqliteProfileRepo) FindByID(id int) (*models.Profile, error) {
	conn := db.GetDB()

	var p models.Profile
	err := conn.QueryRow(`
		SELECT id, user_id, name, age, age_rating, created_at
		FROM profiles
		WHERE id = ?
	`, id).Scan(
		&p.ID,
		&p.UserID,
		&p.Name,
		&p.Age,
		&p.AgeRating,
		&p.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *sqliteProfileRepo) FindByUserID(userID int) ([]models.Profile, error) {
	conn := db.GetDB()

	rows, err := conn.Query(`
		SELECT id, user_id, name, age, age_rating, created_at
		FROM profiles
		WHERE user_id = ?
		ORDER BY id ASC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []models.Profile
	for rows.Next() {
		var p models.Profile
		if err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.Name,
			&p.Age,
			&p.AgeRating,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}
	return profiles, nil
}

func (r *sqliteProfileRepo) CountByUserID(userID int) (int, error) {
	conn := db.GetDB()

	var count int
	err := conn.QueryRow(`
		SELECT COUNT(*)
		FROM profiles
		WHERE user_id = ?
	`, userID).Scan(&count)
	return count, err
}

func (r *sqliteProfileRepo) Delete(id int) error {
	conn := db.GetDB()

	_, err := conn.Exec(`DELETE FROM profiles WHERE id = ?`, id)
	return err
}
