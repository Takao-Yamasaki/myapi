package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu sync.Mutex
)

func newTraceID() int {
	var no int 
	
	mu.Lock()
	no = logNo
	// IDカウンタをインクリメント
	logNo += 1
	mu.Unlock()
	
	return no
}

// コンテキストにトレースIDを付加する
func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, "traceID", traceID)
}

type traceIDKey struct{}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})

	// 型アサーション
	if idInt, ok := id.(int); ok {
		return	idInt
	}
	return 0
}