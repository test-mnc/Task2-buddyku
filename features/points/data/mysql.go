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

func (repo *mysqlPointRepository) SelectPointByIdArticle(idArticle int) (points.Core, error) {
	var point Point

	res := repo.db.Preload("Article").Preload("User").Preload("Company").Where("article_id = ?", idArticle).First(&point)
	if res.Error != nil {
		return points.Core{}, res.Error
	}

	return point.toCore(), nil
}
