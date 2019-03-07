// memo 包提供了一个对类型 Func 并发不安全的函数记忆功能
package memo

// Memo 缓存了调用 Func 的结果
type Memo struct {
	requests chan request
}

// Func 是用于记忆的函数类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// request 是一条请求消息，key 需要用 Func 来调用
type request struct {
	key      string
	response chan<- result // 客户端需要单个 result
}

type entry struct {
	res   result
	ready chan struct{}
}

func (e *entry) call(f Func, key string) {
	// 执行函数
	e.res.value, e.res.err = f(key)
	// 通知数据已准备完毕
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等待该数据准备完毕
	<-e.ready
	// 向客户端发送结果
	response <- e.res
}

// New 返回 f 的函数记忆，客户端之后需要调用 Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// 对这个 key 的第一次请求
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // 调用 f(key)
		}
		go e.deliver(req.response)
	}
}
