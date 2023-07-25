package Q0146_test

import (
	"github.com/kkAtGithub/leetcode/go/LC00/Q0146"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lruCache := Q0146.Constructor(3)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	lruCache.Put(4, 4)

	t.Log(lruCache.Get(4))
	t.Log(lruCache.Get(3))
	t.Log(lruCache.Get(2))
	t.Log(lruCache.Get(1))

	lruCache.Put(5, 5)

	t.Log(lruCache.Get(1))
	t.Log(lruCache.Get(2))
	t.Log(lruCache.Get(3))
	t.Log(lruCache.Get(4))
	t.Log(lruCache.Get(5))
}
