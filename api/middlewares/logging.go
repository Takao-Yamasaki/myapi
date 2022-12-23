package middlewares

import (
	"log"
	"net/http"
)

// 自作のResponseWriterを作る
type resLoggingWriter struct {	
	http.ResponseWriter
	code int
}

// コンストラクタを作る
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// WriteHeaderメソッドを作る
func (rsw *resLoggingWriter) WriteHeader(code int) {
	// resLoggingWriter構造体のcodeフィールドに使うレスポンスコードを保存する
	rsw.code = code
	// HTTPレスポンスに使うレスポンスコードを指定
	// オーバーライドした
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// リクエスト情報をロギング
		log.Println(req.RequestURI, req.Method)

		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Println("res: ", rlw.code)
	})
}
