package tasks

import (
	"fmt"
	"lab2ann/structures"
	"strconv"
	"strings"
)

func BaseballScore(operations string) (int, bool) {
	operations = strings.NewReplacer("[", " ", "]", " ", ",", " ").Replace(operations)
	var stack structures.Stack
	for _, operation := range strings.Fields(operations) {
		var score int
		switch operation {
		case "C":
			if _, ok := stack.Pop(); !ok {
				return 0, false
			}
			continue
		case "D":
			text, ok := stack.Top()
			if !ok {
				return 0, false
			}
			last, _ := strconv.Atoi(text)
			if last > structures.MaxInt32/2 {
				return 0, false
			}
			score = last * 2
		case "+":
			if stack.Length() < 2 {
				return 0, false
			}
			text, _ := stack.Pop()
			previous, _ := stack.Top()
			stack.Push(text)
			first, _ := strconv.Atoi(previous)
			last, _ := strconv.Atoi(text)
			if first > structures.MaxInt32-last {
				return 0, false
			}
			score = first + last
		default:
			value, ok := PositiveInt(operation)
			if !ok {
				return 0, false
			}
			score = value
		}
		stack.Push(strconv.Itoa(score))
	}
	result := 0
	for stack.Length() > 0 {
		text, _ := stack.Pop()
		score, _ := strconv.Atoi(text)
		if result > structures.MaxInt32-score {
			return 0, false
		}
		result += score
	}
	return result, true
}
func (a *App) RunBaseball() {
	fmt.Fprintln(a.out, "Введите операции, например [5, 2, C, D, +]:")
	operations, _ := a.line()
	if result, ok := BaseballScore(operations); ok {
		fmt.Fprintln(a.out, "Сумма очков:", result)
	} else {
		fmt.Fprintln(a.out, "Ошибка: неверная операция, недостаточно очков или переполнение int")
	}
}
