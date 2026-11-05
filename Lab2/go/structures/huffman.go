package structures

import "strings"

type hufNode struct {
	symbol      byte
	weight      int
	left, right *hufNode
}
type Huffman struct {
	root      *hufNode
	codes     HashTable
	frequency [256]int
}

func isLeaf(node *hufNode) bool { return node.left == nil && node.right == nil }
func NewHuffman(text string) *Huffman {
	h := &Huffman{codes: *NewHashTable(8)}
	for i := 0; i < len(text); i++ {
		h.frequency[text[i]]++
	}
	var nodes [256]*hufNode
	count := 0
	for symbol, freq := range h.frequency {
		if freq > 0 {
			nodes[count] = &hufNode{symbol: byte(symbol), weight: freq}
			count++
		}
	}
	for count > 1 {
		min1 := 0
		for i := 1; i < count; i++ {
			if nodes[i].weight < nodes[min1].weight {
				min1 = i
			}
		}
		min2 := -1
		for i := 0; i < count; i++ {
			if i != min1 && (min2 < 0 || nodes[i].weight < nodes[min2].weight) {
				min2 = i
			}
		}
		parent := &hufNode{weight: nodes[min1].weight + nodes[min2].weight, left: nodes[min1], right: nodes[min2]}
		nodes[min1] = parent
		nodes[min2] = nodes[count-1]
		count--
	}
	if count > 0 {
		h.root = nodes[0]
	}
	h.buildCodes(h.root, "")
	return h
}
func (h *Huffman) buildCodes(node *hufNode, code string) {
	if node == nil {
		return
	}
	if isLeaf(node) {
		if code == "" {
			code = "0"
		}
		h.codes.Insert(string([]byte{node.symbol}), code)
		return
	}
	h.buildCodes(node.left, code+"0")
	h.buildCodes(node.right, code+"1")
}
func (h *Huffman) Frequency(symbol byte) int       { return h.frequency[symbol] }
func (h *Huffman) Code(symbol byte) (string, bool) { return h.codes.Get(string([]byte{symbol})) }
func (h *Huffman) Encode(text string) (string, bool) {
	var result strings.Builder
	for i := 0; i < len(text); i++ {
		code, found := h.Code(text[i])
		if !found {
			return "", false
		}
		result.WriteString(code)
	}
	return result.String(), true
}
func (h *Huffman) Decode(bits string) (string, bool) {
	if h.root == nil {
		return "", bits == ""
	}
	var result strings.Builder
	if isLeaf(h.root) {
		for i := 0; i < len(bits); i++ {
			if bits[i] != '0' {
				return "", false
			}
			result.WriteByte(h.root.symbol)
		}
		return result.String(), true
	}
	node := h.root
	for i := 0; i < len(bits); i++ {
		switch bits[i] {
		case '0':
			node = node.left
		case '1':
			node = node.right
		default:
			return "", false
		}
		if isLeaf(node) {
			result.WriteByte(node.symbol)
			node = h.root
		}
	}
	if node != h.root {
		return "", false
	}
	return result.String(), true
}
