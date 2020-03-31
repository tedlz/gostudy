// Package memo 提供并发安全且非阻塞的缓存
// 对不同 key 进行并发请求
// 对相同 key 进行并发请求时，需要等待第一个并发请求完成
// 使用 monitor goroutine 实现
package memo

// !+ Func

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

// !- Func

// !+ get

// 一个 request 是一条消息， Func 是 request 的 key
type request struct {
	key      string
	response chan<- result // client 想要一个结果
}

// Memo *
type Memo struct{ requests chan request }

// New 返回 f 的缓存，客户端必须随后调用 Close
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Get *
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

// Close *
func (memo *Memo) Close() { close(memo.requests) }

// !- get

// !+ monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// 对 key 的首次请求
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // 调用 f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// 评估这个函数
	e.res.value, e.res.err = f(key)
	// 广播 ready 状态
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等待 ready 状态
	<-e.ready
	// 将结果发给客户端
	response <- e.res
}

// !- monitor
