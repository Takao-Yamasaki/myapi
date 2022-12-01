package main

import (
	"log"
	"net/http"

	"github.com/yourname/reponame/handlers"
)

func main() {
	
	// 定義したハンドラをサーバーで使用するように登録
	// パスとハンドラを対応づける
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)

	// サーバ起動時のログを出力する
	log.Println("server start at port 8080")

	// サーバー起動
	// log.Fatal: 重大なエラーが発生した際に、ログを出力させた上で、プログラムを終了させる。
	log.Fatal(http.ListenAndServe(":8080", nil))
}