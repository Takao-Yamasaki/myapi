// ハンドラの定義を記述
package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	// バイトスライスなんらかの形で用意
	// content-lengthフィールドの値を取得
	length ,err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}
	reqBodybuffer := make([]byte, length)

	// リクエストボディを読み出し
	// エラーの内容がEOFか判定するために、errors.ls関数を使用
	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err,io.EOF) {
		// EOF以外だった場合、500番エラーを返却
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	} 

	// ボディをクローズする
	defer req.Body.Close()

	article := models.Article1
	// jsonエンコードする
	jsonData, err := json.Marshal(article)
	if err != nil {
		// エンコードに失敗したということで、500番エラー（サーバー内部でのエラー）を出す
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
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

	articleList := []models.Article{models.Article1, models.Article2}

	jsonData, err := json.Marshal(articleList)

	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// 記事No.xxの投稿データを取得するエンドポイント
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprint("fails to encode json (artcleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// 記事にいいね！をつけるハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

// 記事にコメントを投稿するハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
