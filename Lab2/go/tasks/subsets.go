package tasks

import (
	"fmt"
	"io"
	"lab2ann/structures"
)

func sortSubsets(elements *structures.Array, left, right int) {
	if left >= right {
		return
	}
	middle := elements.Get((left + right) / 2)
	i, j := left, right
	for i <= j {
		for elements.Get(i) < middle {
			i++
		}
		for elements.Get(j) > middle {
			j--
		}
		if i <= j {
			elements.Swap(i, j)
			i++
			j--
		}
	}
	sortSubsets(elements, left, j)
	sortSubsets(elements, i, right)
}
func buildSubsets(out io.Writer, elements, current *structures.Array, start, left int, first *bool) {
	if left == 0 {
		if !*first {
			fmt.Fprint(out, ", ")
		}
		fmt.Fprint(out, "{")
		for i := 0; i < current.Length(); i++ {
			if i > 0 {
				fmt.Fprint(out, ", ")
			}
			fmt.Fprint(out, current.Get(i))
		}
		fmt.Fprint(out, "}")
		*first = false
		return
	}
	for i := start; i <= elements.Length()-left; i++ {
		if i > start && elements.Get(i) == elements.Get(i-1) {
			continue
		}
		current.AddEnd(elements.Get(i))
		buildSubsets(out, elements, current, i+1, left-1, first)
		current.DeleteLast()
	}
}
func PrintAllSubsets(out io.Writer, elements *structures.Array) {
	sortSubsets(elements, 0, elements.Length()-1)
	current := structures.NewArray(elements.Length())
	first := true
	fmt.Fprint(out, "[")
	for size := 0; size <= elements.Length(); size++ {
		buildSubsets(out, elements, current, 0, size, &first)
	}
	fmt.Fprintln(out, "]")
}
func (a *App) RunSubsets() {
	fmt.Fprint(a.out, "Количество элементов массива: ")
	n, ok := a.integer()
	if !ok || n < 0 {
		fmt.Fprintln(a.out, "Ошибка: неверный размер")
		return
	}
	elements := structures.NewArray(n)
	fmt.Fprintln(a.out, "Введите элементы через пробел:")
	for i := 0; i < n; i++ {
		value, ok := a.token()
		if !ok {
			fmt.Fprintln(a.out, "Ошибка: элементов меньше, чем указано")
			return
		}
		elements.AddEnd(value)
	}
	fmt.Fprint(a.out, "Подмножества: ")
	PrintAllSubsets(a.out, elements)
}
