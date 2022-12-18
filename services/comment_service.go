package services

import (
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

// ハンドラ層がComment構造体関連で呼び出したい処理
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
