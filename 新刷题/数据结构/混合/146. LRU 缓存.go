package 混合

import (
	"fmt"
	"strings"
)

type LRUCache struct {
	MyLRU *MyLRU
}

func Constructor(capacity int) LRUCache {
	return LRUCache{MyLRU: NewMyLRU(capacity)}
}

func (c *LRUCache) Get(key int) int {
	return c.MyLRU.Get(key)
}

func (c *LRUCache) Put(key int, value int) {
	c.MyLRU.Put(key, value)
}
