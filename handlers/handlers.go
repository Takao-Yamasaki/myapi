// ハンドラの定義を記述
package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	// クエリパラメータを取得
	queryMap := req.URL.Query()

	var  page int
	// パラメータのpageが一つ以上あるなら、
	// キー:pageの存在確認 true, falseをokに格納
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		// パラメータpageに対応する１つ目の値を取得し、数値に変換する
		var err error
		page , err = strconv.Atoi(p[0])
		// 数値に変換できない場合は400番エラーを返す
		// 400: ユーザーリクエストの値が不正である
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	// パラメータが存在しなかった場合は1を付与する
	} else {
		page = 1
	}
	
	resString := fmt.Sprintf("Article List (page %d)\n", page)
	io.WriteString(w, resString)

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
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)

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
