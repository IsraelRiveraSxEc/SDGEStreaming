// internal/models/profile.go
package models

import "time"

type Profile struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Name      string    `db:"name"`
	Age       int       `db:"age"`
	AgeRating string    `db:"age_rating"`
	CreatedAt time.Time `db:"created_at"`
}
