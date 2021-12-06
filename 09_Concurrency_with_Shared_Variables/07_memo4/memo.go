// Package memo, first try
package memo

import "sync"

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	cache map[string]*entry // zero value of *entry is nil
	mu    sync.Mutex
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}
