package structures

type listNode struct {
	value      string
	next, prev *listNode
}
type List struct{ head, tail *listNode }

func (l *List) addEnd(value string) *listNode {
	node := &listNode{value: value, prev: l.tail}
	if l.tail != nil {
		l.tail.next = node
	} else {
		l.head = node
	}
	l.tail = node
	return node
}
func (l *List) deleteNode(node *listNode) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
	node.next = nil
	node.prev = nil
}
