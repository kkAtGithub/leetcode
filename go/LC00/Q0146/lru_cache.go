package Q0146

type CacheItem struct {
	Key        int
	Val        int
	Prev, Next *CacheItem
}

type LRUCache struct {
	len       int
	cap       int
	head      *CacheItem
	tail      *CacheItem
	kvMapping map[int]*CacheItem
}

func (c *LRUCache) delete(ci *CacheItem) {
	ci.Prev.Next = ci.Next
	ci.Next.Prev = ci.Prev
	ci.Next = nil
	ci.Prev = nil
}

func (c *LRUCache) insert(ci *CacheItem) {
	if c.head.Next != nil {
		ci.Next = c.head.Next
		c.head.Next.Prev = ci
	}
	ci.Prev = c.head
	c.head.Next = ci
}

func (c *LRUCache) RemoveLast() {
	if c.head.Next == c.tail && c.tail.Prev == c.head {
		return
	}
	p := c.tail.Prev
	delete(c.kvMapping, p.Key)
	c.delete(p)
	c.len--
}

func (c *LRUCache) MoveToFirst(ci *CacheItem) {
	c.delete(ci)
	c.insert(ci)
}

func Constructor(capacity int) LRUCache {
	head := &CacheItem{
		Key: -1,
		Val: -1,
	}
	tail := &CacheItem{
		Key: -2,
		Val: -2,
	}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		cap:       capacity,
		head:      head,
		tail:      tail,
		kvMapping: map[int]*CacheItem{},
	}
}

func (c *LRUCache) Get(key int) int {
	if ci, ok := c.kvMapping[key]; ok {
		c.MoveToFirst(ci)
		return ci.Val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if ci, ok := c.kvMapping[key]; ok {
		ci.Val = value
		c.MoveToFirst(ci)
		return
	}
	if c.len+1 > c.cap {
		c.RemoveLast()
	}
	ci := &CacheItem{
		Key: key,
		Val: value,
	}
	c.insert(ci)
	c.kvMapping[ci.Key] = ci
	c.len++
}
