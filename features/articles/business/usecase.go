package business

import "test/mnc/features/articles"

type articleUsecase struct {
	articleData articles.Data
}

func NewArticleBusiness(artcData articles.Data) articles.Business {
	return &articleUsecase{
		articleData: artcData,
	}
}

func (uc *articleUsecase) GetArticleById(idArticle int) (data articles.Core, err error) {
	data, err = uc.articleData.SelectArticleById(idArticle)
	return data, err
}

func (uc *articleUsecase) PostArticle(inputArticle articles.Core) (row int, idArticle int, err error) {
	row, idArticle, err = uc.articleData.InsertArticle(inputArticle)
	return row, idArticle, err
}
