// memo 包提供了一个对类型 Func 并发不安全的函数记忆功能
package memo

import "sync"

// Memo 缓存了调用 Func 的结果
type Memo struct {
	f     Func
	mu    sync.Mutex // 保护 cache
	cache map[string]*entry
}

// Func 是用于记忆的函数类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get 是并发安全的
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// 对 key 的第一次访问，这个 goroutine 负责计算数据和广播数据
		// 已准备完毕的消息
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // 广播数据已准备完毕的消息
	} else {
		// 对这个 key 的重复访问
		memo.mu.Unlock()

		<-e.ready // 等待数据准备完毕
	}

	return e.res.value, e.res.err
}
