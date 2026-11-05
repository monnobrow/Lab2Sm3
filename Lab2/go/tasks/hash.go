package tasks

import (
	"fmt"
	"lab2ann/structures"
	"strconv"
	"strings"
	"time"
)

func measureOpenHash(n int, doubleHash bool) ([4]float64, bool) {
	keys, missing := structures.NewArray(n), structures.NewArray(n)
	for i := 0; i < n; i++ {
		keys.AddEnd(strconv.Itoa(i*37 + 11))
		missing.AddEnd(strconv.Itoa((i+n)*37 + 11))
	}
	table := structures.NewOpenTable(11, doubleHash, false)
	var times [4]float64
	start := time.Now()
	for i := 0; i < n; i++ {
		table.Insert(keys.Get(i), "value")
	}
	times[0] = float64(time.Since(start).Nanoseconds()) / 1000 / float64(n)
	found := 0
	start = time.Now()
	for i := 0; i < n; i++ {
		if _, ok := table.Get(keys.Get(i)); ok {
			found++
		}
	}
	times[1] = float64(time.Since(start).Nanoseconds()) / 1000 / float64(n)
	start = time.Now()
	for i := 0; i < n; i++ {
		if _, ok := table.Get(missing.Get(i)); ok {
			found++
		}
	}
	times[2] = float64(time.Since(start).Nanoseconds()) / 1000 / float64(n)
	start = time.Now()
	for i := 0; i < n; i++ {
		table.Delete(keys.Get(i))
	}
	times[3] = float64(time.Since(start).Nanoseconds()) / 1000 / float64(n)
	return times, found == n && table.Length() == 0
}
func (a *App) benchmarkOpenHash(n int) {
	linear, linearOK := measureOpenHash(n, false)
	doubleHash, doubleOK := measureOpenHash(n, true)
	if !linearOK || !doubleOK {
		fmt.Fprintln(a.out, "Ошибка: таблица работает неправильно")
	}
	names := [4]string{"Добавление", "Поиск существующего", "Поиск отсутствующего", "Удаление"}
	fmt.Fprintf(a.out, "\nN = %d, микросекунды на операцию\nОперация | Линейное | Двойное\n", n)
	for i, name := range names {
		fmt.Fprintf(a.out, "%s | %.3f | %.3f\n", name, linear[i], doubleHash[i])
	}
}
func (a *App) demoFolding() {
	fmt.Fprint(a.out, "Введите числовой ключ: ")
	key, _ := a.line()
	sum, formula, ok := structures.FoldingSum(key)
	if !ok {
		fmt.Fprintln(a.out, "Ошибка: нужен числовой ключ с суммой частей в пределах int")
		return
	}
	fmt.Fprintf(a.out, "%s = %d\n", formula, sum)
	linear := structures.NewOpenTable(11, false, true)
	doubleHash := structures.NewOpenTable(11, true, true)
	linear.Insert(key, strconv.Itoa(sum))
	doubleHash.Insert(key, strconv.Itoa(sum))
	value, _ := linear.Get(key)
	fmt.Fprintf(a.out, "Линейное: найдено %s, начальная ячейка %d\n", value, structures.FoldHash(key, linear.Capacity()))
	value, _ = doubleHash.Get(key)
	fmt.Fprintf(a.out, "Двойное: найдено %s, начальная ячейка %d\n", value, structures.FoldHash(key, doubleHash.Capacity()))
}
func (a *App) demoOpenTables() {
	fmt.Fprintln(a.out, "Введите числовые ключи через пробел:")
	text, _ := a.line()
	linear := structures.NewOpenTable(11, false, true)
	doubleHash := structures.NewOpenTable(11, true, true)
	for _, key := range strings.Fields(text) {
		sum, _, ok := structures.FoldingSum(key)
		if !ok {
			fmt.Fprintln(a.out, "Пропущен неверный ключ:", key)
			continue
		}
		linear.Insert(key, strconv.Itoa(sum))
		doubleHash.Insert(key, strconv.Itoa(sum))
	}
	fmt.Fprintln(a.out, "Линейное пробирование:")
	linear.Print(a.out)
	fmt.Fprintln(a.out, "Двойное хеширование:")
	doubleHash.Print(a.out)
}
func (a *App) RunHashTask() {
	fmt.Fprint(a.out, "1 — метод свёртки, 2 — сравнение таблиц, 3 — показать коллизии: ")
	choice, ok := a.integer()
	if !ok {
		return
	}
	a.line()
	switch choice {
	case 1:
		a.demoFolding()
	case 2:
		a.benchmarkOpenHash(1000)
		a.benchmarkOpenHash(5000)
		a.benchmarkOpenHash(10000)
	case 3:
		a.demoOpenTables()
	default:
		fmt.Fprintln(a.out, "Нет такого режима")
	}
}
