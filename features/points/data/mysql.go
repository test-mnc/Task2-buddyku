package data

import (
	"test/mnc/features/points"

	"gorm.io/gorm"
)

type mysqlPointRepository struct {
	db *gorm.DB
}

func NewPointRepository(conn *gorm.DB) points.Data {
	return &mysqlPointRepository{
		db: conn,
	}
}

func (repo *mysqlPointRepository) InsertPoint(inputValue points.Core) (int, error) {
	var point = fromCore(inputValue)

	res := repo.db.Create(&point)
	if res.Error != nil {
		return 0, res.Error
	}

	return 1, nil
}
