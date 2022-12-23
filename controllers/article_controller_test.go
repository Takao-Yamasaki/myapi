package controllers_test

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/yourname/reponame/controllers"
	"github.com/yourname/reponame/services"

	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

// テストしたい関数で使うリソースを作成

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("DB setup fail")
		os.Exit(1)
	}
	
	ser := services.NewMyAppService(db)
	aCon = controllers.NewArticleController(ser)
	
	m.Run()
}



func TestArticleListHandler(t *testing.T) {
	// 2.テスト対象の関数に入れるinputを定義
	var tests = []struct {
		name string
		query string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ハンドラに渡す２つの引数
			// w http.ResponseWriter, req *http.Requestを用意する
			url := fmt.Sprintf("http://localhost:8080/article/list?page=%s, tt.query")
			req := httptest.NewRequest(http.MethodGet, url, nil)
			res := httptest.NewRecorder()

			// 3.テスト対象の関数を実行してoutputを得る
			aCon.ArticleListHandler(res, req)

			// 4.outputが期待通りかチェック
			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: StatusCode: want %d but %d\n, tt.resultCode, res.Code")
			}
		})
	}
}
