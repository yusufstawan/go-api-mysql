package ageRatingCategory

import (
	"api-mysql/config"
	"api-mysql/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "age_rating_category"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll AgeRatingCategory
func GetAll(ctx context.Context) ([]models.AgeRatingCategory, error) {

	var ratings []models.AgeRatingCategory

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var rating models.AgeRatingCategory
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&rating.ID,
			&rating.Name,
			&rating.Description,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		rating.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		rating.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		ratings = append(ratings, rating)
	}

	return ratings, nil
}

// Insert AgeRatingCategory
func Insert(ctx context.Context, rating models.AgeRatingCategory) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name, description, created_at, updated_at) values('%v','%v', NOW(), NOW())", table,
		rating.Name,
		rating.Description)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update AgeRatingCategory
func Update(ctx context.Context, rating models.AgeRatingCategory, id string) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name ='%s', description = '%s', updated_at = NOW() where id = %s",
		table,
		rating.Name,
		rating.Description,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete AgeRatingCategory
func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
