// Package day2 — переменные, константы, iota, базовые типы и приведение типов.
package day2

// Weekday — день недели. Monday = 0, Sunday = 6.
type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var weekdayNames = [...]string{"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}

// String возвращает название дня по-русски, для неизвестного значения — "Weekday(N)".
func (d Weekday) String() string {
	// TODO: ваше решение
	return ""
}

// IsWeekend сообщает, выходной ли это день.
func (d Weekday) IsWeekend() bool {
	// TODO: ваше решение
	return false
}

// Размеры в байтах: KB = 1024, MB = 1024*KB, GB = 1024*MB.
const (
	KB = 0 // TODO: через iota и сдвиг <<
	MB = 0 // TODO
	GB = 0 // TODO
)

// FormatBytes форматирует размер: меньше KB — "N B", иначе с одним знаком после точки
// в самой крупной подходящей единице: "1.5 KB", "10.0 MB", "3.5 GB".
func FormatBytes(n int64) string {
	// TODO: ваше решение
	return ""
}

// CelsiusToFahrenheit переводит градусы Цельсия в Фаренгейты: F = C*9/5 + 32.
func CelsiusToFahrenheit(c float64) float64 {
	// TODO: ваше решение
	return 0
}
