package request

import "test/mnc/features/points"

type Point struct {
	Value     string `json:"value"`
	ArticleID int    `json:"article_id"`
	UserID    int    `json:"user_id"`
	CompanyID int    `json:"company_id"`
}

func ToCore(req Point) points.Core {
	return points.Core{
		Value: req.Value,
		Article: points.Article{
			ID: req.ArticleID,
		},
		Company: points.Company{
			ID: req.CompanyID,
		},
		User: points.User{
			ID: req.UserID,
		},
	}
}
