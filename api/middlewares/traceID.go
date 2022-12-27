package middlewares

import (
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