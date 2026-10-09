// Package day1 — первое упражнение: модуль, пакет, тесты.
package day1

import (
	"strings"
)

// Greet возвращает приветствие вида "Привет, <name>!".
// Пробелы по краям имени отбрасываются; если имя пустое, вместо него "мир".
func Greet(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Привет, мир!"
	}
	return "Привет, " + name + "!"
}
