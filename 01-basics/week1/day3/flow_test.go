package day3

import (
	"errors"
	"slices"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if got := FizzBuzz(15); !slices.Equal(got, want) {
		t.Errorf("FizzBuzz(15) =\n%q\nwant\n%q", got, want)
	}
	if got := FizzBuzz(0); len(got) != 0 {
		t.Errorf("FizzBuzz(0) = %q, want пустой слайс", got)
	}
}

func TestCollatzSteps(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 0},
		{2, 1},
		{6, 8},
		{27, 111},
	}
	for _, tt := range tests {
		got, err := CollatzSteps(tt.n)
		if err != nil {
			t.Errorf("CollatzSteps(%d): неожиданная ошибка %v", tt.n, err)
			continue
		}
		if got != tt.want {
			t.Errorf("CollatzSteps(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
	for _, n := range []int{0, -5} {
		if _, err := CollatzSteps(n); !errors.Is(err, ErrNotPositive) {
			t.Errorf("CollatzSteps(%d): err = %v, want ErrNotPositive", n, err)
		}
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{100, "отлично"},
		{90, "отлично"},
		{89, "хорошо"},
		{75, "хорошо"},
		{74, "удовлетворительно"},
		{60, "удовлетворительно"},
		{59, "неудовлетворительно"},
		{0, "неудовлетворительно"},
		{-1, "некорректная оценка"},
		{101, "некорректная оценка"},
	}
	for _, tt := range tests {
		if got := Grade(tt.score); got != tt.want {
			t.Errorf("Grade(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}

func TestDeferOrder(t *testing.T) {
	want := []int{3, 2, 1}
	if got := DeferOrder(); !slices.Equal(got, want) {
		t.Errorf("DeferOrder() = %v, want %v", got, want)
	}
}
