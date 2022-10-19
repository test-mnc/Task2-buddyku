package presentation

import (
	"net/http"
	"strconv"
	"test/mnc/features/articles"
	_requestArticle "test/mnc/features/articles/presentation/request"
	_responseArticle "test/mnc/features/articles/presentation/response"
	_helper "test/mnc/helpers"
	"test/mnc/middlewares"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	articleBusiness articles.Business
}

func NewArticleHandler(business articles.Business) *ArticleHandler {
	return &ArticleHandler{
		articleBusiness: business,
	}
}

func (h *ArticleHandler) ReadArticleById(c echo.Context) error {
	id := c.Param("idArticle")
	idArtic, _ := strconv.Atoi(id)
	res, err := h.articleBusiness.GetArticleById(idArtic)
	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to get data article"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("success to get data article", _responseArticle.FromCore(res)))
}

func (h *ArticleHandler) PostingArticle(c echo.Context) error {
	idUser, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("invalid token"))
	}
	var article = _requestArticle.Article{
		UserID: idUser,
	}

	errBind := c.Bind(&article)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to insert bind data"))
	}

	res, idArticle, err := h.articleBusiness.PostArticle(_requestArticle.ToCore(article))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))
	}

	if res == 0 {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert/record data"))
	}

	var dataId = map[string]interface{}{
		"id_article": idArticle,
	}

	return c.JSON(http.StatusOK, _helper.ResponseSuccesWithData("success to insert data", dataId))
}
