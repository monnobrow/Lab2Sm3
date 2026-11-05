package tasks

import (
	"fmt"
	"lab2ann/structures"
	"strconv"
	"strings"
)

func ReadPartitionSet(text string, set *structures.Set) bool {
	text = strings.NewReplacer("{", " ", "}", " ", ",", " ").Replace(text)
	for _, item := range strings.Fields(text) {
		value, ok := PositiveInt(item)
		if !ok {
			return false
		}
		set.Add(strconv.Itoa(value))
	}
	return set.Length() > 0
}
func number(elements *structures.Array, index int) int {
	value, _ := strconv.Atoi(elements.Get(index))
	return value
}
func sortPartition(elements *structures.Array, left, right int) {
	if left >= right {
		return
	}
	middle := number(elements, (left+right)/2)
	i, j := left, right
	for i <= j {
		for number(elements, i) > middle {
			i++
		}
		for number(elements, j) < middle {
			j--
		}
		if i <= j {
			elements.Swap(i, j)
			i++
			j--
		}
	}
	sortPartition(elements, left, j)
	sortPartition(elements, i, right)
}
func fillEqualGroups(elements *structures.Array, index, target int, sums, groups []int) bool {
	if index == elements.Length() {
		return true
	}
	value := number(elements, index)
	previousSum := -1
	for group := range sums {
		if sums[group] == previousSum {
			continue
		}
		previousSum = sums[group]
		if value > target-sums[group] {
			continue
		}
		sums[group] += value
		groups[index] = group
		if fillEqualGroups(elements, index+1, target, sums, groups) {
			return true
		}
		sums[group] -= value
		groups[index] = -1
		if sums[group] == 0 {
			break
		}
	}
	return false
}
func EqualPartition(elements *structures.Array, target int) ([]int, int, bool) {
	if target < 1 || elements.Length() == 0 {
		return nil, 0, false
	}
	total := 0
	for i := 0; i < elements.Length(); i++ {
		value, ok := PositiveInt(elements.Get(i))
		if !ok || value > target || total > structures.MaxInt32-value {
			return nil, 0, false
		}
		total += value
	}
	if total%target != 0 {
		return nil, 0, false
	}
	count := total / target
	if count > elements.Length() {
		return nil, 0, false
	}
	sortPartition(elements, 0, elements.Length()-1)
	groups := make([]int, elements.Length())
	for i := range groups {
		groups[i] = -1
	}
	if !fillEqualGroups(elements, 0, target, make([]int, count), groups) {
		return nil, count, false
	}
	return groups, count, true
}
func (a *App) RunPartition() {
	fmt.Fprintln(a.out, "Введите множество, например {4,10,5,1,3,7}:")
	text, _ := a.line()
	set := structures.NewSet()
	if !ReadPartitionSet(text, set) {
		fmt.Fprintln(a.out, "Ошибка: нужно непустое множество натуральных чисел")
		return
	}
	fmt.Fprint(a.out, "Сумма каждого подмножества: ")
	target, ok := a.integer()
	if !ok || target < 1 {
		fmt.Fprintln(a.out, "Ошибка: нужна положительная сумма")
		return
	}
	elements := set.Elements()
	groups, count, found := EqualPartition(elements, target)
	if !found {
		fmt.Fprintln(a.out, "Ошибка: множество нельзя разбить на подмножества с заданной суммой")
		return
	}
	for group := 0; group < count; group++ {
		fmt.Fprintf(a.out, "Подмножество %d: {", group+1)
		first := true
		for i := 0; i < elements.Length(); i++ {
			if groups[i] == group {
				if !first {
					fmt.Fprint(a.out, ", ")
				}
				fmt.Fprint(a.out, elements.Get(i))
				first = false
			}
		}
		fmt.Fprintf(a.out, "}, сумма: %d\n", target)
	}
}
