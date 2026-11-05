package tasks

import (
	"fmt"
	"lab2ann/structures"
	"strings"
)

func (a *App) RunLRU() {
	fmt.Fprint(a.out, "Ёмкость кэша и количество запросов: ")
	capacity, ok := a.integer()
	if !ok {
		fmt.Fprintln(a.out, "Ошибка: нужны неотрицательные числа")
		return
	}
	q, ok := a.integer()
	if !ok || capacity < 0 || q < 0 {
		fmt.Fprintln(a.out, "Ошибка: нужны неотрицательные числа")
		return
	}
	cache := structures.NewLRUCache(capacity)
	fmt.Fprintln(a.out, "Введите запросы SET ключ значение или GET ключ:")
	var result strings.Builder
	for i := 0; i < q; i++ {
		command, ok := a.token()
		if !ok {
			fmt.Fprintln(a.out, "Ошибка: не хватает запросов")
			return
		}
		key, ok := a.token()
		if !ok {
			fmt.Fprintln(a.out, "Ошибка: не хватает запросов")
			return
		}
		switch command {
		case "SET":
			value, ok := a.token()
			if !ok {
				fmt.Fprintln(a.out, "Ошибка: нужно значение")
				return
			}
			cache.Set(key, value)
		case "GET":
			if result.Len() > 0 {
				result.WriteByte(' ')
			}
			result.WriteString(cache.Get(key))
		default:
			fmt.Fprintln(a.out, "Ошибка: неизвестная команда")
			return
		}
	}
	fmt.Fprintln(a.out, "Результат:", result.String())
}
