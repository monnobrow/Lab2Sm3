package structures

import (
	"fmt"
	"io"
)

type hashNode struct {
	key, value string
	next       *hashNode
}
type HashTable struct {
	data []*hashNode
	size int
}

func NewHashTable(capacity int) *HashTable {
	if capacity < 1 {
		capacity = 1
	}
	return &HashTable{data: make([]*hashNode, capacity)}
}

// uint32 даёт такое же переполнение хеша, как unsigned int в C++-версии.
func Hash(key string, capacity int) int {
	var value uint32
	for i := 0; i < len(key); i++ {
		value = value*31 + uint32(key[i])
	}
	return int(value % uint32(capacity))
}
func (h *HashTable) resize() {
	data := make([]*hashNode, max(1, len(h.data)*2))
	for _, head := range h.data {
		for node := head; node != nil; {
			next := node.next
			index := Hash(node.key, len(data))
			node.next = data[index]
			data[index] = node
			node = next
		}
	}
	h.data = data
}

// Insert возвращает true для нового ключа и false при обновлении значения.
func (h *HashTable) Insert(key, value string) bool {
	if len(h.data) == 0 {
		h.data = make([]*hashNode, 8)
	}
	index := Hash(key, len(h.data))
	for node := h.data[index]; node != nil; node = node.next {
		if node.key == key {
			node.value = value
			return false
		}
	}
	h.data[index] = &hashNode{key, value, h.data[index]}
	h.size++
	if h.size >= len(h.data)-len(h.data)/4 {
		h.resize()
	}
	return true
}
func (h *HashTable) Get(key string) (string, bool) {
	if len(h.data) == 0 {
		return "", false
	}
	for node := h.data[Hash(key, len(h.data))]; node != nil; node = node.next {
		if node.key == key {
			return node.value, true
		}
	}
	return "", false
}
func (h *HashTable) Contains(key string) bool { _, found := h.Get(key); return found }
func (h *HashTable) Delete(key string) bool {
	if len(h.data) == 0 {
		return false
	}
	index := Hash(key, len(h.data))
	var previous *hashNode
	node := h.data[index]
	for node != nil && node.key != key {
		previous = node
		node = node.next
	}
	if node == nil {
		return false
	}
	if previous == nil {
		h.data[index] = node.next
	} else {
		previous.next = node.next
	}
	h.size--
	return true
}
func (h *HashTable) Length() int { return h.size }
func (h *HashTable) Keys() *Array {
	keys := NewArray(h.size)
	for _, head := range h.data {
		for node := head; node != nil; node = node.next {
			keys.AddEnd(node.key)
		}
	}
	return keys
}

type Set struct{ table HashTable }

func NewSet() *Set                        { return &Set{table: *NewHashTable(8)} }
func (s *Set) Add(value string) bool      { return s.table.Insert(value, "") }
func (s *Set) Delete(value string) bool   { return s.table.Delete(value) }
func (s *Set) Contains(value string) bool { return s.table.Contains(value) }
func (s *Set) Length() int                { return s.table.Length() }
func (s *Set) Elements() *Array           { return s.table.Keys() }
func (s *Set) Print(out io.Writer) {
	keys := s.Elements()
	fmt.Fprint(out, "{")
	for i := 0; i < keys.Length(); i++ {
		if i > 0 {
			fmt.Fprint(out, ", ")
		}
		fmt.Fprint(out, keys.Get(i))
	}
	fmt.Fprintln(out, "}")
}
