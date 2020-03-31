// Package memo 向 type Func 提供并发安全的缓存
// 对不同 key 进行并发请求
// 对相同 key 进行并发请求会导致重复读取 / 写入
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
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// 在两个关键部分之间，有几个 goroutines 可能竞相计算 f(key) 并更新 map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
