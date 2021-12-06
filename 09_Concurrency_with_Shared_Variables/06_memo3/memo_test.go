package memo_test

import (
	memo "The.Go.Programming.Language/09_Concurrency_with_Shared_Variables/06_memo3"
	"The.Go.Programming.Language/09_Concurrency_with_Shared_Variables/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
