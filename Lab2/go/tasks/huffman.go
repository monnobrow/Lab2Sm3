package tasks

import (
	"fmt"
	"lab2ann/structures"
)

func (a *App) RunHuffman() {
	fmt.Fprintln(a.out, "Введите строку:")
	text, _ := a.line()
	if text == "" {
		fmt.Fprintln(a.out, "Строка пустая")
		return
	}
	h := structures.NewHuffman(text)
	fmt.Fprintln(a.out, "Символ — частота — код:")
	for i := 0; i < 256; i++ {
		freq := h.Frequency(byte(i))
		if freq == 0 {
			continue
		}
		code, _ := h.Code(byte(i))
		if i == ' ' {
			fmt.Fprint(a.out, "пробел")
		} else if i > 32 && i < 127 {
			fmt.Fprintf(a.out, "%c", byte(i))
		} else {
			fmt.Fprint(a.out, "байт ", i)
		}
		fmt.Fprintf(a.out, " — %d — %s\n", freq, code)
	}
	bits, _ := h.Encode(text)
	decoded, _ := h.Decode(bits)
	fmt.Fprintln(a.out, "Код:", bits)
	fmt.Fprintln(a.out, "Расшифровка:", decoded)
	fmt.Fprintln(a.out, "Введите код для расшифровки по этому дереву (пустая строка — пропустить):")
	bits, ok := a.line()
	if ok && bits != "" {
		if result, correct := h.Decode(bits); correct {
			fmt.Fprintln(a.out, "Результат:", result)
		} else {
			fmt.Fprintln(a.out, "Ошибка: код содержит не 0/1 или обрывается внутри символа")
		}
	}
}
