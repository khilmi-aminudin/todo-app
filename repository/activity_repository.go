package repository

import (
	"database/sql"
	"fmt"

	"github.com/khilmi-aminudin/todo-app/models"
	_ "github.com/lib/pq"
)

type ActivityRepository interface {
	Get(activityParam models.Activity) models.Activity
	GetAll() []models.Activity
	Create(activityParam models.Activity) models.Activity
	Update(activityParam models.Activity) models.Activity
	Delete(activityParam models.Activity) error
}

type activityRepository struct {
	db *sql.DB
}

var queryString string

func NewActivityRepository(db *sql.DB) ActivityRepository {
	return &activityRepository{
		db: db,
	}
}

func (repo *activityRepository) Get(activityParam models.Activity) models.Activity {
	queryString = "SELECT * FROM activity WHERE id = $1 LIMIT 1;"

	row := repo.db.QueryRow(queryString, activityParam.ID)

	var activity models.Activity

	row.Scan(
		&activity.ID,
		&activity.Email,
		&activity.Title,
		&activity.CreatedAt,
		&activity.UpdatedAt,
		&activity.DeletedAt,
	)

	return activity

}
func (repo *activityRepository) GetAll() []models.Activity {
	queryString = "SELECT * FROM activity;"

	rows, err := repo.db.Query(queryString)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err.Error()))
	}

	var activities []models.Activity

	for rows.Next() {
		var a models.Activity
		rows.Scan(
			&a.ID,
			&a.Email,
			&a.Title,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.DeletedAt,
		)
		activities = append(activities, a)

	}

	return activities
}

func (repo *activityRepository) Create(activityParam models.Activity) models.Activity {
	queryString = `
	INSERT INTO activity (
		email,
		title,
	) VALUES (
		$1, $2
	) RETURNING 
	id, email, title, created_at, updated_at, deleted_at;`

	_, err := repo.db.Exec(queryString, activityParam.Email, activityParam.Title, activityParam.CreatedAt, activityParam.CreatedAt, activityParam.UpdatedAt)

	if err != nil {
		panic(fmt.Sprintf("error: %v", err.Error()))
	}

	activity := models.Activity{
		Email:     activityParam.Email,
		Title:     activityParam.Title,
		CreatedAt: activityParam.CreatedAt,
		UpdatedAt: activityParam.UpdatedAt,
		DeletedAt: activityParam.DeletedAt,
	}
	return activity
}

func (repo *activityRepository) Update(activityParam models.Activity) models.Activity {
	queryString = `UPDATE activity SET email = $2, title = $3, updated_at = $4 WHERE id = $1 RETURNING id, email, title, created_at, updated_at, deleted_at;`

	row := repo.db.QueryRow(queryString, activityParam.ID, activityParam.Email, activityParam.Title, activityParam.UpdatedAt)

	var activity models.Activity

	err := row.Scan(&activity.ID, &activity.Email, &activity.Title, &activity.CreatedAt, &activity.UpdatedAt, &activity.DeletedAt)

	if err != nil {
		panic(err)
	}
	return activity
}
func (repo *activityRepository) Delete(activityParam models.Activity) error {
	queryString = `DELETE FROM activity WHERE id = $1;`

	_, err := repo.db.Exec(queryString, activityParam.ID)
	return err
}
