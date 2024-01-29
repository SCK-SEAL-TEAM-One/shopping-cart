package point

import (
	"github.com/jmoiron/sqlx"
)

type PointRepository interface {
	GetPoints(userID int) ([]Point, error)
	CreatePoint(userID int, amount int) (int, error)
}

type PointRepositoryMySQL struct {
	DBConnection *sqlx.DB
}

func (repository PointRepositoryMySQL) GetPoints(userID int) ([]Point, error) {
	var Points []Point
	err := repository.DBConnection.Select(&Points, `
		SELECT id, user_id, amount
		FROM points
		WHERE  user_id = ?
	`, userID)
	return Points, err
}

func (repository PointRepositoryMySQL) CreatePoint(userID int, amount int) (int, error) {
	sqlResult := repository.DBConnection.MustExec("INSERT INTO points (user_id, amount) VALUE (?,?)", userID, amount)
	insertedId, err := sqlResult.LastInsertId()
	return int(insertedId), err
}
