package structures

// Array повторяет массив строк из ЛР1: size — занятые ячейки, data — вся ёмкость.
type Array struct {
	data []string
	size int
}

func NewArray(capacity int) *Array {
	if capacity < 1 {
		capacity = 1
	}
	return &Array{data: make([]string, capacity)}
}
func (a *Array) AddEnd(value string) {
	if a.size == len(a.data) {
		replacement := make([]string, max(1, len(a.data)*2))
		copy(replacement, a.data)
		a.data = replacement
	}
	a.data[a.size] = value
	a.size++
}
func (a *Array) Get(index int) string {
	if index < 0 || index >= a.size {
		panic("индекс вне массива")
	}
	return a.data[index]
}
func (a *Array) Swap(i, j int) {
	if i < 0 || j < 0 || i >= a.size || j >= a.size {
		panic("индекс вне массива")
	}
	a.data[i], a.data[j] = a.data[j], a.data[i]
}
func (a *Array) DeleteLast() bool {
	if a.size == 0 {
		return false
	}
	a.size--
	a.data[a.size] = ""
	return true
}
func (a *Array) Length() int { return a.size }
