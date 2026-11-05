package tasks

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type App struct {
	in       *bufio.Reader
	out, err io.Writer
}

func NewApp(in io.Reader, out, err io.Writer) *App {
	return &App{in: bufio.NewReader(in), out: out, err: err}
}
func (a *App) line() (string, bool) {
	text, err := a.in.ReadString('\n')
	if err != nil && len(text) == 0 {
		return "", false
	}
	return strings.TrimSuffix(text, "\n"), true
}
func (a *App) token() (string, bool) {
	var text string
	_, err := fmt.Fscan(a.in, &text)
	return text, err == nil
}
func (a *App) integer() (int, bool) {
	text, ok := a.token()
	if !ok {
		return 0, false
	}
	number, err := strconv.ParseInt(text, 10, 32)
	return int(number), err == nil
}
func (a *App) RunTask(number int) int {
	switch number {
	case 1:
		a.RunBaseball()
	case 2:
		a.RunSetTask()
	case 3:
		a.RunPartition()
	case 4:
		a.RunSubsets()
	case 5:
		a.RunHuffman()
	case 6:
		a.RunHashTask()
	case 7:
		a.RunLRU()
	default:
		fmt.Fprintln(a.out, "Нет такого задания")
		return 1
	}
	return 0
}
func (a *App) menu() {
	fmt.Fprint(a.out, "\nLab2 Ann — номер 19\n1 — Бейсбол со стеком\n2 — АТД Множество\n3 — Разбиение множества\n4 — Все подмножества массива\n5 — Хаффман\n6 — Линейное и двойное хеширование, свёртка\n7 — LRU-кэш\n0 — Выход\n> ")
}
func Run(args []string, in io.Reader, out, err io.Writer) int {
	a := NewApp(in, out, err)
	if len(args) == 1 && args[0] == "--help" {
		fmt.Fprint(out, "lab2 — меню\nlab2 1..7 — отдельное задание\nlab2 --file set.txt --query 'SETADD 5' — команда множества\n")
		return 0
	}
	if len(args) == 1 && len(args[0]) == 1 && args[0][0] >= '1' && args[0][0] <= '7' {
		return a.RunTask(int(args[0][0] - '0'))
	}
	if len(args) > 0 {
		return a.RunSetArgs(args)
	}
	for {
		a.menu()
		number, ok := a.integer()
		if !ok {
			break
		}
		a.line()
		if number == 0 {
			break
		}
		a.RunTask(number)
	}
	return 0
}
func PositiveInt(text string) (int, bool) {
	if text == "" {
		return 0, false
	}
	value := 0
	for i := 0; i < len(text); i++ {
		if text[i] < '0' || text[i] > '9' {
			return 0, false
		}
		digit := int(text[i] - '0')
		if value > (2147483647-digit)/10 {
			return 0, false
		}
		value = value*10 + digit
	}
	return value, value > 0
}
