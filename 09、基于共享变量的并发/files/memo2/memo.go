// Package memo 向 type Func 提供并发安全的缓存
// 并发请求由 Mutex 序列化
package memo

import "sync"

// Memo 会缓存 Func 的结果
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

// Func 是要缓存的类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New *
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get 是线程安全的
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
