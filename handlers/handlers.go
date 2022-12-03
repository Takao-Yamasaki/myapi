// ハンドラの定義を記述
package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourname/reponame/models"
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
	//jsonをデコード(json→構造体)
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// jsonエンコードする(構造体→json)
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// ブログ一覧を取得するハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// クエリパラメータを取得
	queryMap := req.URL.Query()

	var page int
	// パラメータのpageが一つ以上あるなら、
	// キー:pageの存在確認 true, falseをokに格納
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		// パラメータpageに対応する１つ目の値を取得し、数値に変換する
		var err error
		page, err = strconv.Atoi(p[0])
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

	log.Println(page)

	// jsonエンコードする（構造体→json）
	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

// 記事No.xxの投稿データを取得するエンドポイント
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	log.Println(articleID)

	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

// 記事にいいね！をつけるハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// jsonデコードする（json→構造体）
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// jsonエンコードする（構造体→json）
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// 記事にコメントを投稿するハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// jsonデコードする（json→構造体）
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// jsonエンコードする（構造体→json）
	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
