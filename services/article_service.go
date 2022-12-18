package services

// ハンドラ層がarticle構造体関連で呼び出したい処理

import (
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	// 記事の詳細を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// コメント一覧を取得
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// コメント一覧をArticle構造体に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		return models.Article{}, err
	}

	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
