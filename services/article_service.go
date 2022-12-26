package services

// ハンドラ層がarticle構造体関連で呼び出したい処理

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/yourname/reponame/apperrors"
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	// 待ち合わせ機構で、内部カウンタは0
	var wg sync.WaitGroup
	// 内部カウンタを+2
	wg.Add(2)

	var amu sync.Mutex
	var cmu sync.Mutex

	// 記事の詳細を取得
	go func(db *sql.DB, articleID int) {
		// 内部カウンタを-1
		defer wg.Done()
		newarticle, err := repositories.SelectArticleDetail(db, articleID)
		amu.Lock()
		article, articleGetErr = newarticle, err
		amu.Unlock()
	}(s.db, articleID)

	// コメント一覧を取得
	go func(db *sql.DB, articleID int) {
		// 内部カウンタを-1
		defer wg.Done()
		newcommentList, err := repositories.SelectCommentList(db, articleID)
		cmu.Lock()
		commentList, commentGetErr = newcommentList, err
		cmu.Unlock()
	}(s.db, articleID)

	// カウンタが0になるまで待機
	wg.Wait()

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	// コメント一覧をArticle構造体に紐づける
	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update nice count")
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
