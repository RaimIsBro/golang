// Package day4 — функции: несколько результатов, вариадические аргументы,
// функции как значения и замыкания.
package day4

// Sum возвращает сумму всех аргументов; без аргументов — 0.
func Sum(nums ...int) int {
	// TODO: ваше решение
	return 0
}

// MinMax возвращает минимум и максимум; ok = false, если аргументов нет.
func MinMax(nums ...int) (lo, hi int, ok bool) {
	// TODO: ваше решение
	return 0, 0, false
}

// Map возвращает новый слайс, где к каждому элементу применена f. Вход не меняется.
func Map(nums []int, f func(int) int) []int {
	// TODO: ваше решение
	return nil
}

// Filter возвращает новый слайс из элементов, для которых keep вернула true.
func Filter(nums []int, keep func(int) bool) []int {
	// TODO: ваше решение
	return nil
}

// NewCounter возвращает функцию, которая при каждом вызове возвращает 1, 2, 3, ...
// Каждый счётчик независим.
func NewCounter() func() int {
	// TODO: ваше решение
	return func() int { return 0 }
}

// Compose возвращает функцию x -> f(g(x)).
func Compose(f, g func(int) int) func(int) int {
	// TODO: ваше решение
	return nil
}
