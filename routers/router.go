package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/controllers"
)

func NewRouter(con *controllers.MyAppController) *mux.Router {
	// ルータrを明示的に宣言
	r := mux.NewRouter()

	// 定義したハンドラをサーバーで使用するように登録
	// パスとハンドラを対応づける
	r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

	return r
}
