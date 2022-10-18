package data

import (
	"test/mnc/features/articles"

	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) articles.Data {
	return &mysqlArticleRepository{
		db: conn,
	}
}

func (repo *mysqlArticleRepository) SelectArticleById(idArticle int) (articles.Core, error) {
	var dataArticle Article
	res := repo.db.Preload("User").First(&dataArticle, idArticle)
	if res.Error != nil {
		return articles.Core{}, res.Error
	}

	return dataArticle.toCore(), nil
}

func (repo *mysqlArticleRepository) InsertArticle(input articles.Core) (int, int, error) {
	var article = fromCore(input)

	res := repo.db.Create(&article)
	if res.Error != nil {
		return 0, 0, res.Error
	}

	return 1, int(article.ID), nil
}
