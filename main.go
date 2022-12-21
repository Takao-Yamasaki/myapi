package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yourname/reponame/api"
)

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	// サーバ起動時のログを出力する
	log.Println("server start at port 8080")

	_, err0 := strconv.Atoi("a")
	fmt.Printf("err0: [%T] %v\n", err0, err0)

	err1 := errors.Unwrap(err0)
	fmt.Printf("err1: [%T] %v\n", err1, err1)

	err2 := errors.Unwrap(err1)
	fmt.Printf("err2: [%T] %v\n", err2, err2)
	
	// サーバー起動
	// log.Fatal: 重大なエラーが発生した際に、ログを出力させた上で、プログラムを終了させる。
	// 第二引数には、ルーターを指定する
	log.Fatal(http.ListenAndServe(":8080", r))
}
