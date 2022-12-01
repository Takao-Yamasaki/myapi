package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// ハンドラ: HTTPリクエストを受け取って、それに対するHTTPレスポンスの内容をコネクションに書き込む
	// 引数にhttp.ResponseWriter型とhttp.Request型をとる。引数はなし。

	// ハンドラの定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// req: http.Requestで受けとって、w: http.ResponseWriterに書き込む
		io.WriteString(w, "Hello, world!\n")
	}

	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Article ... \n")
	}

	// ブログ一覧を取得するハンドラ
	articleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}

	// 記事No.xxの投稿データを取得するエンドポイント
	articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article No.1\n")
	}

	// 記事にいいね！をつけるハンドラ
	postNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}

	// 記事にコメントを投稿するハンドラ
	postCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	// 定義したハンドラをサーバーで使用するように登録
	// パスとハンドラを対応づける
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", articleListHandler)
	http.HandleFunc("/article/1", articleDetailHandler)
	http.HandleFunc("/article/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	// サーバ起動時のログを出力する
	log.Println("server start at port 8080")

	// サーバー起動
	// log.Fatal: 重大なエラーが発生した際に、ログを出力させた上で、プログラムを終了させる。
	log.Fatal(http.ListenAndServe(":8080", nil))
}