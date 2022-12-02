// ハンドラの定義を記述
package handlers

import (
	"io"
	"net/http"
)

// ハンドラ: HTTPリクエストを受け取って、それに対するHTTPレスポンスの内容をコネクションに書き込む
// 引数にhttp.ResponseWriter型とhttp.Request型をとる。引数はなし。

// ハンドラの定義
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world!\n")
	
	/*
	// req: http.Requestで受けとって、w: http.ResponseWriterに書き込む
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")
	} else {
		// Invalid methodを405番のステータスコードとともに返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
	*/
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article ... \n")
	
	/*if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article ... \n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
	*/
}

// ブログ一覧を取得するハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
	
	/*
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
	*/
}

// 記事No.xxの投稿データを取得するエンドポイント
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article No.1\n")
	
	/*
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article No.1\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
	*/
}

// 記事にいいね！をつけるハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
	
	/*
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
	*/
}

// 記事にコメントを投稿するハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
	
	/*
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
	*/
}
