package structures

type stackNode struct {
	value string
	next  *stackNode
}
type Stack struct {
	head *stackNode
	size int
}

func (s *Stack) Push(value string) { s.head = &stackNode{value, s.head}; s.size++ }
func (s *Stack) Top() (string, bool) {
	if s.head == nil {
		return "", false
	}
	return s.head.value, true
}
func (s *Stack) Pop() (string, bool) {
	if s.head == nil {
		return "", false
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}
func (s *Stack) Length() int { return s.size }
