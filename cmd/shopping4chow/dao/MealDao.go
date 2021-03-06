package dao

import (
	"shopping4chow/cmd/shopping4chow/models"

	"github.com/jackc/pgx/v4"
)

type MealDao interface {
	GetMeal(conn *pgx.Conn, findMeal models.Meal) []models.Meal
	RemoveMeal(id int) error
	GetAllMeals() []models.Meal
	AddMeal(userName string, Meal models.Meal) (int, error)
}
