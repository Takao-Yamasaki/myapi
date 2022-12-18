package services

import (
	"github.com/yourname/reponame/models"
	"github.com/yourname/reponame/repositories"
)

// ハンドラ層がComment構造体関連で呼び出したい処理
func PostCommentService(comment models.Comment) (models.Comment, error) {
	// データベースに接続
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, nil
	}

	return newComment, nil
}
