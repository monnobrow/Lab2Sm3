package structures

import "strconv"

type lruEntry struct {
	value string
	node  *listNode
}
type LRUCache struct {
	capacity, size int
	index          HashTable
	keys           List
	entries        []lruEntry
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity < 0 {
		capacity = 0
	}
	return &LRUCache{capacity: capacity, index: *NewHashTable(8), entries: make([]lruEntry, capacity)}
}
func (c *LRUCache) touch(slot int) {
	key := c.entries[slot].node.value
	c.keys.deleteNode(c.entries[slot].node)
	c.entries[slot].node = c.keys.addEnd(key)
}
func (c *LRUCache) Get(key string) string {
	saved, found := c.index.Get(key)
	if !found {
		return "-1"
	}
	slot, _ := strconv.Atoi(saved)
	c.touch(slot)
	return c.entries[slot].value
}
func (c *LRUCache) Set(key, value string) {
	if c.capacity == 0 {
		return
	}
	if saved, found := c.index.Get(key); found {
		slot, _ := strconv.Atoi(saved)
		c.entries[slot].value = value
		c.touch(slot)
		return
	}
	var slot int
	if c.size == c.capacity {
		victim := c.keys.head.value
		saved, _ := c.index.Get(victim)
		slot, _ = strconv.Atoi(saved)
		c.keys.deleteNode(c.keys.head)
		c.index.Delete(victim)
	} else {
		slot = c.size
		c.size++
	}
	c.entries[slot] = lruEntry{value: value, node: c.keys.addEnd(key)}
	c.index.Insert(key, strconv.Itoa(slot))
}
func (c *LRUCache) Length() int { return c.size }
