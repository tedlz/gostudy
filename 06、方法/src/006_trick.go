package main

import "sync"

// 006、通过嵌入结构体来扩展类型
func main() {

}

var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

// LookUp *
func LookUp(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

// 下面这个版本在功能上是一致的，但将两个包级别的变量放在了 cache 这个 struct 内

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

// LookUp2 *
func LookUp2(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
