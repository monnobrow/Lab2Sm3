package structures

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestArrayStackAndSet(t *testing.T) {
	a := NewArray(0)
	for i := 0; i < 100; i++ {
		a.AddEnd(strconv.Itoa(i))
	}
	if a.Length() != 100 || a.Get(99) != "99" {
		t.Fatal("array growth")
	}
	a.Swap(0, 99)
	if a.Get(0) != "99" {
		t.Fatal("array swap")
	}
	for a.DeleteLast() {
	}
	if a.Length() != 0 {
		t.Fatal("array pop")
	}
	var zero Array
	zero.AddEnd("zero")
	if zero.Get(0) != "zero" {
		t.Fatal("zero array")
	}
	var stack Stack
	if _, ok := stack.Top(); ok {
		t.Fatal("empty top")
	}
	if _, ok := stack.Pop(); ok {
		t.Fatal("empty pop")
	}
	for _, v := range []string{"A", "B", "C"} {
		stack.Push(v)
	}
	for _, want := range []string{"C", "B", "A"} {
		got, ok := stack.Pop()
		if !ok || got != want {
			t.Fatal("LIFO", got, want)
		}
	}
	if stack.Length() != 0 {
		t.Fatal("stack size")
	}
	set := NewSet()
	if !set.Add("x") || set.Add("x") || !set.Contains("x") || set.Length() != 1 {
		t.Fatal("set uniqueness")
	}
	if !set.Delete("x") || set.Delete("x") || set.Contains("x") {
		t.Fatal("set removal")
	}
	var out bytes.Buffer
	set.Print(&out)
	if out.String() != "{}\n" {
		t.Fatal(out.String())
	}
}
func TestHashTablesAgainstMap(t *testing.T) {
	tables := []interface {
		Insert(string, string) bool
		Get(string) (string, bool)
		Delete(string) bool
		Length() int
	}{
		NewHashTable(1), NewOpenTable(3, false, false), NewOpenTable(3, true, false), NewOpenTable(3, false, true), NewOpenTable(3, true, true),
	}
	for kind, table := range tables {
		model := map[string]string{}
		random := rand.New(rand.NewSource(19))
		for step := 0; step < 4000; step++ {
			key := strconv.Itoa(random.Intn(200))
			value := strconv.Itoa(step)
			if random.Intn(3) != 0 {
				_, existed := model[key]
				if table.Insert(key, value) == existed {
					t.Fatalf("table %d insert flag", kind)
				}
				model[key] = value
			} else {
				_, existed := model[key]
				if table.Delete(key) != existed {
					t.Fatalf("table %d delete flag", kind)
				}
				delete(model, key)
			}
			if table.Length() != len(model) {
				t.Fatalf("table %d size", kind)
			}
			for probe := 0; probe < 201; probe++ {
				k := strconv.Itoa(probe)
				want, found := model[k]
				got, ok := table.Get(k)
				if ok != found || (ok && got != want) {
					t.Fatalf("table %d step %d key %s got %q/%v want %q/%v", kind, step, k, got, ok, want, found)
				}
			}
		}
		table.Insert("", "")
		got, ok := table.Get("")
		if !ok || got != "" {
			t.Fatal("empty key/value")
		}
	}
	// Коллизии должны сохраняться при удалении головы цепочки или создании надгробия.
	for _, doubleHash := range []bool{false, true} {
		table := NewOpenTable(11, doubleHash, true)
		for _, key := range []string{"11", "22", "33", "44"} {
			table.Insert(key, key)
		}
		table.Delete("11")
		if got, ok := table.Get("44"); !ok || got != "44" {
			t.Fatal("collision after tombstone")
		}
		table.Insert("55", "55")
		var out bytes.Buffer
		table.Print(&out)
		if !strings.Contains(out.String(), "55 -> 55") {
			t.Fatal("print table")
		}
	}
	var zero HashTable
	zero.Insert("a", "b")
	if got, _ := zero.Get("a"); got != "b" {
		t.Fatal("zero hash")
	}
	var open OpenTable
	open.Insert("a", "b")
	if got, _ := open.Get("a"); got != "b" {
		t.Fatal("zero open table")
	}
}
func TestFolding(t *testing.T) {
	for _, c := range []struct {
		key, formula string
		sum          int
	}{{"523456795", "523+456+795", 1774}, {"1234", "123+4", 127}, {"000012", "000+012", 12}, {"0", "0", 0}} {
		got, formula, ok := FoldingSum(c.key)
		if !ok || got != c.sum || formula != c.formula {
			t.Fatal(c, got, formula, ok)
		}
		if FoldHash(c.key, 11) != c.sum%11 {
			t.Fatal("fold index")
		}
	}
	for _, key := range []string{"", "-1", "a", "12 34", "12x"} {
		if _, _, ok := FoldingSum(key); ok {
			t.Fatal("accepted invalid key", key)
		}
	}
	if FoldHash("abc", 11) != Hash("abc", 11) {
		t.Fatal("fallback hash")
	}
	if _, _, ok := FoldingSum(strings.Repeat("999", MaxInt32/999+1)); ok {
		t.Fatal("folding overflow")
	}
}
func TestLRUAgainstModel(t *testing.T) {
	for capacity := 0; capacity <= 7; capacity++ {
		cache := NewLRUCache(capacity)
		values := map[string]string{}
		order := []string{}
		random := rand.New(rand.NewSource(19))
		touch := func(key string) {
			for i, k := range order {
				if k == key {
					order = append(order[:i], order[i+1:]...)
					break
				}
			}
			order = append(order, key)
		}
		for step := 0; step < 4000; step++ {
			key := strconv.Itoa(random.Intn(12))
			if random.Intn(3) != 0 {
				value := strconv.Itoa(step)
				cache.Set(key, value)
				if capacity > 0 {
					if _, ok := values[key]; !ok && len(values) == capacity {
						delete(values, order[0])
						order = order[1:]
					}
					values[key] = value
					touch(key)
				}
			} else {
				want, found := values[key]
				if !found {
					want = "-1"
				} else {
					touch(key)
				}
				if got := cache.Get(key); got != want {
					t.Fatalf("capacity %d step %d got %s want %s", capacity, step, got, want)
				}
			}
			if cache.Length() != len(values) {
				t.Fatal("LRU size")
			}
			node := cache.keys.head
			var previous *listNode
			for _, k := range order {
				if node == nil || node.value != k || node.prev != previous {
					t.Fatal("LRU list order")
				}
				previous = node
				node = node.next
			}
			if node != nil || cache.keys.tail != previous {
				t.Fatal("LRU tail")
			}
		}
	}
	if NewLRUCache(-1).Get("x") != "-1" {
		t.Fatal("negative capacity")
	}
}
func TestHuffmanRoundtrip(t *testing.T) {
	random := rand.New(rand.NewSource(19))
	all := make([]byte, 256)
	for i := range all {
		all[i] = byte(i)
	}
	cases := []string{"", "aaaa", "abracadabra", "Анна Петрова", string(all), "a\x00b"}
	for n := 0; n < 200; n++ {
		data := make([]byte, random.Intn(300))
		for i := range data {
			data[i] = byte(random.Intn(256))
		}
		cases = append(cases, string(data))
	}
	for _, text := range cases {
		h := NewHuffman(text)
		bits, ok := h.Encode(text)
		if !ok {
			t.Fatal("encode")
		}
		decoded, ok := h.Decode(bits)
		if !ok || decoded != text {
			t.Fatalf("roundtrip %q", text)
		}
		if _, ok := h.Decode("not bits"); ok {
			t.Fatal("invalid bits")
		}
	}
	h := NewHuffman("abracadabra")
	bits, _ := h.Encode("abracadabra")
	if bits != "01101001110011110110100" {
		t.Fatal("different tie-breaking", bits)
	}
	if _, ok := h.Decode("111"); ok {
		t.Fatal("incomplete symbol")
	}
	if _, ok := NewHuffman("a").Decode("1"); ok {
		t.Fatal("single symbol invalid bit")
	}
	if _, ok := NewHuffman("a").Encode("b"); ok {
		t.Fatal("unknown symbol")
	}
	for i := 0; i < 256; i++ {
		code, ok := h.Code(byte(i))
		if ok && h.Frequency(byte(i)) <= 0 {
			t.Fatal(fmt.Sprint(code))
		}
	}
}
