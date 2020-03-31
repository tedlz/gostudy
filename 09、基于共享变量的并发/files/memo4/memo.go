// Package memo 提供并发安全的缓存
// 对不同 key 进行并发请求
// 对相同 key 进行并发请求时，需要等待第一个并发请求完成
// 使用互斥锁实现
package memo

import "sync"

// Memo 会缓存 Func 的结果
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

// Func 是要缓存的类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // ready 后关闭
}

// New *
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get 是线程安全的
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// 对 key 的首次请求
		// 这个 goroutine 负责计算值并广播 ready 状态
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // 广播 ready 状态
	} else {
		// 对 key 的重复请求
		memo.mu.Unlock()
		<-e.ready // 等待 ready 状态
	}
	return e.res.value, e.res.err
}
