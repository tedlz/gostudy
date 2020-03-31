// Package memo 向 Func 提供并发不安全的缓存
package memo

// Memo 会缓存 Func 的结果
type Memo struct {
	f     Func
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

// Get *
// 注意：非线程安全
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
