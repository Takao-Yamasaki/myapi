package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/handlers"
)

func main() {
	// ルータrを明示的に宣言
	r := mux.NewRouter()
	
	// 定義したハンドラをサーバーで使用するように登録
	// パスとハンドラを対応づける
	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバ起動時のログを出力する
	log.Println("server start at port 8080")

	// サーバー起動
	// log.Fatal: 重大なエラーが発生した際に、ログを出力させた上で、プログラムを終了させる。
	// 第二引数には、ルーターを指定する
	log.Fatal(http.ListenAndServe(":8080", r))
}