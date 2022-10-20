package business

import (
	"test/mnc/features/articles"
	"test/mnc/features/points"
)

type articleUsecase struct {
	articleData articles.Data
	pointData   points.Data
}

func NewArticleBusiness(artcData articles.Data, pntData points.Data) articles.Business {
	return &articleUsecase{
		articleData: artcData,
		pointData:   pntData,
	}
}

func (uc *articleUsecase) GetArticleById(idArticle int) (data articles.Core, err error) {
	data, err = uc.articleData.SelectArticleById(idArticle)
	return data, err
}

func (uc *articleUsecase) PostArticle(inputArticle articles.Core) (row int, idArticle int, err error) {
	row, idArticle, err = uc.articleData.InsertArticle(inputArticle)
	if row == 1 {
		row1, idArticle1, err1 := uc.pointData.FirstPoint(inputArticle.User.ID, idArticle, 1, 1)
		return row1, idArticle1, err1
	}
	return row, idArticle, err
}

func (uc *articleUsecase) GetAllArticles() (data []articles.Core, err error) {
	data, err = uc.articleData.SelectAllArticles()
	return data, err
}
