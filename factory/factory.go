package factory

import (
	_userBusiness "test/mnc/features/users/business"
	_userData "test/mnc/features/users/data"
	_userPresentation "test/mnc/features/users/presentation"

	_articleBusiness "test/mnc/features/articles/business"
	_articleData "test/mnc/features/articles/data"
	_articlePresentation "test/mnc/features/articles/presentation"

	_companyBusiness "test/mnc/features/company/business"
	_companyData "test/mnc/features/company/data"
	_companyPresentation "test/mnc/features/company/presentation"

	_pointBusiness "test/mnc/features/points/business"
	_pointData "test/mnc/features/points/data"
	_pointPresentation "test/mnc/features/points/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *_userPresentation.UserHandler
	ArticlePresenter *_articlePresentation.ArticleHandler
	CompanyPresenter *_companyPresentation.CompanyHandler
	PointPresenter   *_pointPresentation.PointHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	UserPresentation := _userPresentation.NewUserHandler(userBusiness)

	articleData := _articleData.NewArticleRepository(dbConn)
	articleBusiness := _articleBusiness.NewArticleBusiness(articleData)
	ArticlePresentation := _articlePresentation.NewArticleHandler(articleBusiness)

	companyData := _companyData.NewCompanyRepository(dbConn)
	companyBusiness := _companyBusiness.NewCompanyBusiness(companyData)
	CompanyPresentation := _companyPresentation.NewCompanyHandler(companyBusiness)

	pointData := _pointData.NewPointRepository(dbConn)
	pointBusiness := _pointBusiness.NewPointBusiness(pointData)
	PointPresentation := _pointPresentation.NewPointHandler(pointBusiness)

	return Presenter{
		UserPresenter:    UserPresentation,
		ArticlePresenter: ArticlePresentation,
		CompanyPresenter: CompanyPresentation,
		PointPresenter:   PointPresentation,
	}
}
