// Package day3 — управление потоком: for, if, switch, defer.
package day3

import "errors"

// ErrNotPositive возвращается, когда аргумент должен быть положительным.
var ErrNotPositive = errors.New("число должно быть положительным")

// FizzBuzz возвращает строки для чисел от 1 до n: "Fizz" для кратных 3,
// "Buzz" для кратных 5, "FizzBuzz" для кратных 15, иначе само число.
func FizzBuzz(n int) []string {
	// TODO: ваше решение
	return nil
}

// CollatzSteps считает шаги гипотезы Коллатца до 1: чётное n делим на 2,
// нечётное заменяем на 3n+1. Для n <= 0 возвращает ErrNotPositive.
func CollatzSteps(n int) (int, error) {
	// TODO: ваше решение
	return 0, nil
}

// Grade переводит баллы 0..100 в оценку: 90+ "отлично", 75+ "хорошо",
// 60+ "удовлетворительно", ниже "неудовлетворительно", вне диапазона "некорректная оценка".
func Grade(score int) string {
	// TODO: ваше решение
	return ""
}

// DeferOrder должна вернуть []int{3, 2, 1}, заполнив слайс только через defer,
// по одному defer на числа 1, 2, 3 в таком порядке.
func DeferOrder() (res []int) {
	// TODO: ваше решение
	return nil
}
