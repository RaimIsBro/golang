// Package day1 — первое упражнение: модуль, пакет, тесты.
package day1

import (
	"fmt"
	"strings"
)

// Greet возвращает приветствие вида "Привет, <name>!".
// Пробелы по краям имени отбрасываются; если имя пустое, вместо него "мир".
func Greet(name string) string {
	if name == "" {
		return "Привет, мир!"
	}
	name = strings.TrimSpace(name)
	return fmt.Sprintf("Привет, %s!", name)
}
