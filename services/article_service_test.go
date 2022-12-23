// ベンチマークテストの作成
package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/yourname/reponame/services"

	_ "github.com/go-sql-driver/mysql"
)

var aSer *services.MyAppService

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err !=  nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	aSer = services.NewMyAppService(db)

	// 個別のベンチマークテストの実行
	m.Run()
}

// ベンチマークテストの実装
// 測りたい関数・メソッドを複数回実行して、その平均を求める
func BenchmarkGetArticleService(b *testing.B) {
	// 引数の準備
	articleID := 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}