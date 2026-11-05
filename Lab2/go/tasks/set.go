package tasks

import (
	"bufio"
	"fmt"
	"io"
	"lab2ann/structures"
	"os"
	"strings"
)

func LoadSet(set *structures.Set, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		value := strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
		if value != "" {
			set.Add(value)
		}
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}
func SaveSet(set *structures.Set, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	elements := set.Elements()
	for i := 0; i < elements.Length(); i++ {
		if _, err = fmt.Fprintln(writer, elements.Get(i)); err != nil {
			file.Close()
			return err
		}
	}
	err = writer.Flush()
	closeErr := file.Close()
	if err != nil {
		return err
	}
	return closeErr
}
func (a *App) RunSetQuery(path, query string) int {
	args := strings.Fields(query)
	command := ""
	if len(args) > 0 {
		command = args[0]
	}
	if command != "SETADD" && command != "SETDEL" && command != "SET_AT" && command != "SETPRINT" {
		fmt.Fprintln(a.err, "Ошибка: неизвестная команда", command)
		return 1
	}
	if command != "SETPRINT" && len(args) < 2 {
		fmt.Fprintln(a.err, "Ошибка: нужно значение элемента")
		return 1
	}
	expected := 2
	if command == "SETPRINT" {
		expected = 1
	}
	if len(args) > expected {
		fmt.Fprintln(a.err, "Ошибка: слишком много аргументов команды")
		return 1
	}
	set := structures.NewSet()
	if LoadSet(set, path) != nil {
		fmt.Fprintln(a.err, "Ошибка: не удалось прочитать файл", path)
		return 1
	}
	value := ""
	if len(args) > 1 {
		value = args[1]
	}
	switch command {
	case "SETADD":
		if !set.Add(value) {
			fmt.Fprintln(a.out, value, "уже есть в множестве")
		} else if SaveSet(set, path) != nil {
			fmt.Fprintln(a.err, "Ошибка: не удалось записать файл", path)
			return 1
		} else {
			fmt.Fprintln(a.out, "->", value)
		}
	case "SETDEL":
		if !set.Delete(value) {
			fmt.Fprintln(a.out, value, "нет в множестве")
		} else if SaveSet(set, path) != nil {
			fmt.Fprintln(a.err, "Ошибка: не удалось записать файл", path)
			return 1
		} else {
			fmt.Fprintln(a.out, "->", value, "удалён")
		}
	case "SET_AT":
		if set.Contains(value) {
			fmt.Fprintln(a.out, "-> TRUE")
		} else {
			fmt.Fprintln(a.out, "-> FALSE")
		}
	case "SETPRINT":
		set.Print(a.out)
	}
	return 0
}
func (a *App) RunSetArgs(args []string) int {
	var path, query string
	for i := 0; i < len(args); i++ {
		if (args[i] == "--file" || args[i] == "--query") && i+1 < len(args) {
			if args[i] == "--file" {
				path = args[i+1]
			} else {
				query = args[i+1]
			}
			i++
		} else {
			fmt.Fprintln(a.err, "Ошибка: неизвестный или неполный параметр", args[i])
			return 1
		}
	}
	if path == "" || query == "" {
		fmt.Fprintln(a.err, "Запуск: lab2 --file set.txt --query 'SETADD x | SETDEL x | SET_AT x'")
		return 1
	}
	return a.RunSetQuery(path, query)
}
func (a *App) RunSetTask() {
	fmt.Fprint(a.out, "Файл множества: ")
	var path string
	for {
		line, ok := a.line()
		if !ok {
			break
		}
		path = strings.TrimLeft(line, " \t\r\v\f")
		if path != "" {
			break
		}
	}
	fmt.Fprint(a.out, "Команда (SETADD x / SETDEL x / SET_AT x / SETPRINT): ")
	query, _ := a.line()
	a.RunSetQuery(path, query)
}
