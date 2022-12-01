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

	// 定義したハンドラをサーバーで使用するように登録
	http.HandleFunc("/", helloHandler)

	// サーバ起動時のログを出力する
	log.Println("server start at port 8080")

	// サーバー起動
	// log.Fatal: 重大なエラーが発生した際に、ログを出力させた上で、プログラムを終了させる。
	log.Fatal(http.ListenAndServe(":8080", nil))
}