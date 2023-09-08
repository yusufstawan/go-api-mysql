package models

import (
	"time"
)

type (
	// Movie
	Movie struct {
		ID                  int       `json:"id"`
		Title               string    `json:"title"`
		Year                int       `json:"year"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
		AgeRatingCategoryID int       `json:"age_rating_category_id"`
	}
)
