package structures

import (
	"fmt"
	"io"
	"strings"
)

const MaxInt32 = 2147483647
const (
	empty = iota
	occupied
	deleted
)

type openCell struct {
	key, value string
	state      int
}
type OpenTable struct {
	data                []openCell
	size, deletedCount  int
	doubleHash, folding bool
}

func isPrime(value int) bool {
	if value < 2 {
		return false
	}
	for i := 2; i <= value/i; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}
func nextPrime(value int) int {
	if value < 3 {
		value = 3
	}
	for !isPrime(value) {
		value++
	}
	return value
}
func NewOpenTable(capacity int, doubleHash, folding bool) *OpenTable {
	return &OpenTable{data: make([]openCell, nextPrime(capacity)), doubleHash: doubleHash, folding: folding}
}
func FoldHash(key string, capacity int) int {
	result := 0
	for i := 0; i < len(key); i += 3 {
		part := 0
		for j := i; j < i+3 && j < len(key); j++ {
			if key[j] < '0' || key[j] > '9' {
				return Hash(key, capacity)
			}
			part = part*10 + int(key[j]-'0')
		}
		result = (result + part) % capacity
	}
	return result
}
func FoldingSum(key string) (int, string, bool) {
	if key == "" {
		return 0, "", false
	}
	sum := 0
	var formula strings.Builder
	for i := 0; i < len(key); i += 3 {
		part := 0
		end := min(i+3, len(key))
		for j := i; j < end; j++ {
			if key[j] < '0' || key[j] > '9' {
				return 0, "", false
			}
			part = part*10 + int(key[j]-'0')
		}
		if sum > MaxInt32-part {
			return 0, "", false
		}
		sum += part
		if i > 0 {
			formula.WriteByte('+')
		}
		formula.WriteString(key[i:end])
	}
	return sum, formula.String(), true
}
func (h *OpenTable) start(key string) int {
	if h.folding {
		return FoldHash(key, len(h.data))
	}
	return Hash(key, len(h.data))
}
func (h *OpenTable) step(key string) int {
	if !h.doubleHash {
		return 1
	}
	value := uint32(2166136261)
	for i := 0; i < len(key); i++ {
		value ^= uint32(key[i])
		value *= 16777619
	}
	return 1 + int(value%uint32(len(h.data)-1))
}
func (h *OpenTable) find(key string) int {
	if len(h.data) == 0 {
		return -1
	}
	index, step := h.start(key), h.step(key)
	for i := 0; i < len(h.data); i++ {
		if h.data[index].state == empty {
			return -1
		}
		if h.data[index].state == occupied && h.data[index].key == key {
			return index
		}
		index = (index + step) % len(h.data)
	}
	return -1
}
func (h *OpenTable) rehash(capacity int) {
	replacement := NewOpenTable(capacity, h.doubleHash, h.folding)
	for _, cell := range h.data {
		if cell.state != occupied {
			continue
		}
		index, step := replacement.start(cell.key), replacement.step(cell.key)
		for replacement.data[index].state == occupied {
			index = (index + step) % len(replacement.data)
		}
		replacement.data[index] = cell
		replacement.size++
	}
	*h = *replacement
}
func (h *OpenTable) Insert(key, value string) bool {
	if len(h.data) == 0 {
		*h = *NewOpenTable(11, h.doubleHash, h.folding)
	}
	if index := h.find(key); index >= 0 {
		h.data[index].value = value
		return false
	}
	if h.size+h.deletedCount+1 >= len(h.data)-len(h.data)/4 {
		capacity := len(h.data)
		if h.size+1 >= capacity/2 {
			capacity *= 2
		}
		h.rehash(capacity)
	}
	index, step := h.start(key), h.step(key)
	for h.data[index].state == occupied {
		index = (index + step) % len(h.data)
	}
	if h.data[index].state == deleted {
		h.deletedCount--
	}
	h.data[index] = openCell{key, value, occupied}
	h.size++
	return true
}
func (h *OpenTable) Get(key string) (string, bool) {
	index := h.find(key)
	if index < 0 {
		return "", false
	}
	return h.data[index].value, true
}
func (h *OpenTable) Delete(key string) bool {
	index := h.find(key)
	if index < 0 {
		return false
	}
	h.data[index] = openCell{state: deleted}
	h.size--
	h.deletedCount++
	return true
}
func (h *OpenTable) Length() int   { return h.size }
func (h *OpenTable) Capacity() int { return len(h.data) }
func (h *OpenTable) Print(out io.Writer) {
	for i, cell := range h.data {
		fmt.Fprintf(out, "[%d] ", i)
		switch cell.state {
		case occupied:
			fmt.Fprintf(out, "%s -> %s", cell.key, cell.value)
		case deleted:
			fmt.Fprint(out, "удалено")
		default:
			fmt.Fprint(out, "-")
		}
		fmt.Fprintln(out)
	}
}
