package migration

import (
	_mArticles "test/mnc/features/articles/data"
	_mCompany "test/mnc/features/company/data"
	_mPoints "test/mnc/features/points/data"
	_mUsers "test/mnc/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mArticles.Article{})
	db.AutoMigrate(&_mCompany.Company{})
	db.AutoMigrate(&_mPoints.Point{})
}
