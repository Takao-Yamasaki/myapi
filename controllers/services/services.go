package services

import "github.com/yourname/reponame/models"

// article関連のメソッド
type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)

}

// comment関連のメソッド
type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}