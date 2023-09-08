package main

import (
	"api-mysql/ageRatingCategory"
	"api-mysql/models"
	"api-mysql/movie"
	"api-mysql/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/movie", GetMovie)
	router.POST("/movie/create", PostMovie)
	router.PUT("/movie/:id/update", UpdateMovie)
	router.DELETE("/movie/:id/delete", DeleteMovie)

	router.GET("/age-rating-category", GetAgeRatingCategory)
	router.POST("/age-rating-category/create", PostAgeRatingCategory)
	router.PUT("/age-rating-category/:id/update", UpdateAgeRatingCategory)
	router.DELETE("/age-rating-category/:id/delete", DeleteAgeRatingCategory)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

// Read
// GetMovie
func GetMovie(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	movies, err := movie.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, movies, http.StatusOK)
}

// Create
// PostMovie
func PostMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mov models.Movie
	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := movie.Insert(ctx, mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateMovie
func UpdateMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var mov models.Movie

	if err := json.NewDecoder(r.Body).Decode(&mov); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idMovie = ps.ByName("id")

	if err := movie.Update(ctx, mov, idMovie); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteMovie
func DeleteMovie(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idMovie = ps.ByName("id")

	if err := movie.Delete(ctx, idMovie); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

// Read
// GetAgeRatingCategory
func GetAgeRatingCategory(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	ratings, err := ageRatingCategory.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, ratings, http.StatusOK)
}

// Create
// PostAgeRatingCategory
func PostAgeRatingCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var rating models.AgeRatingCategory

	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := ageRatingCategory.Insert(ctx, rating); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateAgeRatingCategory
func UpdateAgeRatingCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var rating models.AgeRatingCategory

	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idRating = ps.ByName("id")

	if err := ageRatingCategory.Update(ctx, rating, idRating); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
}

// Delete
// DeleteAgeRatingCategory
func DeleteAgeRatingCategory(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idRating = ps.ByName("id")

	if err := ageRatingCategory.Delete(ctx, idRating); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}
