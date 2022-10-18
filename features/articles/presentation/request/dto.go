package request

import "test/mnc/features/articles"

type Article struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserID  int    `json:"user_id" form:"user_id"`
}

func ToCore(req Article) articles.Core {
	return articles.Core{
		Title:   req.Title,
		Content: req.Content,
		User: articles.User{
			ID: req.UserID,
		},
	}
}
