package routes

import (
	"test/mnc/factory"
	"test/mnc/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.POST("users", presenter.UserPresenter.AddUser)
	e.POST("login", presenter.UserPresenter.Login)
	e.GET("users", presenter.UserPresenter.GetUser, middlewares.JWTMiddleware())

	e.GET("articles/:idArticle", presenter.ArticlePresenter.ReadArticleById)
	e.POST("articles", presenter.ArticlePresenter.PostingArticle, middlewares.JWTMiddleware())
	e.GET("articles", presenter.ArticlePresenter.ReadAllArticles)

	e.POST("company/login", presenter.CompanyPresenter.Login)
	e.POST("company", presenter.CompanyPresenter.AddEmployee)
	e.GET("company/users", presenter.UserPresenter.GetAllUsers, middlewares.JWTMiddleware())
	e.POST("company/points", presenter.PointPresenter.InsertValue, middlewares.JWTMiddleware())

	e.GET("points/:idArticle", presenter.PointPresenter.SelectPointPerArticle, middlewares.JWTMiddleware())
	e.GET("points/users/:idUser", presenter.PointPresenter.SelectPointPerUser, middlewares.JWTMiddleware())
	e.GET("points", presenter.PointPresenter.SelectAllPoints, middlewares.JWTMiddleware())

	return e
}
