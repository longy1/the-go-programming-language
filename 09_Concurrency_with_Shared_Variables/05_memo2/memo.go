// Package memo, first try
package memo

import "sync"

type Memo struct {
	f     Func
	cache map[string]result
	mu    sync.Mutex
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// Get is concurrency-safe, but high-cost
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	defer memo.mu.Unlock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}
