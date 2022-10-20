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

func (repo *mysqlPointRepository) SelectPointByIdArticle(idArticle int) ([]points.Core, error) {
	var point []Point

	res := repo.db.Preload("Article").Preload("User").Preload("Company").Where("article_id = ?", idArticle).Find(&point)
	if res.Error != nil {
		return nil, res.Error
	}

	return toCoreList(point), nil
}

func (repo *mysqlPointRepository) SelectPointByIdUser(idUser int) ([]points.Core, error) {
	var point []Point

	res := repo.db.Preload("User").Preload("Article").Preload("Company").Where("user_id = ?", idUser).Find(&point)

	if res.Error != nil {
		return nil, res.Error
	}

	return toCoreList(point), nil
}

func (repo *mysqlPointRepository) FirstPoint(idUser int, ArticleID int, CompanyID int, value float64) (row int, idArticle int, err error) {
	var point = Point{
		Value:     value,
		UserID:    idUser,
		ArticleID: ArticleID,
		CompanyID: CompanyID,
	}

	result := repo.db.Create(&point)
	if result.Error != nil {
		return 0, 0, result.Error
	}
	return 1, point.ArticleID, nil
}
