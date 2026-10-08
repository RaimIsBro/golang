package day4

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d, want 0", got)
	}
	if got := Sum(1, 2, 3); got != 6 {
		t.Errorf("Sum(1, 2, 3) = %d, want 6", got)
	}
	nums := []int{10, -4, 5}
	if got := Sum(nums...); got != 11 {
		t.Errorf("Sum(nums...) = %d, want 11", got)
	}
}

func TestMinMax(t *testing.T) {
	lo, hi, ok := MinMax(3, -2, 9, 0)
	if !ok || lo != -2 || hi != 9 {
		t.Errorf("MinMax(3, -2, 9, 0) = %d, %d, %v; want -2, 9, true", lo, hi, ok)
	}
	lo, hi, ok = MinMax(7)
	if !ok || lo != 7 || hi != 7 {
		t.Errorf("MinMax(7) = %d, %d, %v; want 7, 7, true", lo, hi, ok)
	}
	if _, _, ok = MinMax(); ok {
		t.Error("MinMax() вернула ok = true для пустого списка")
	}
}

func TestMap(t *testing.T) {
	double := func(x int) int { return x * 2 }
	if got := Map([]int{1, 2, 3}, double); !slices.Equal(got, []int{2, 4, 6}) {
		t.Errorf("Map(double) = %v, want [2 4 6]", got)
	}
	in := []int{1, 2, 3}
	Map(in, double)
	if !slices.Equal(in, []int{1, 2, 3}) {
		t.Errorf("Map изменила входной слайс: %v", in)
	}
}

func TestFilter(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	if got := Filter([]int{1, 2, 3, 4, 5, 6}, even); !slices.Equal(got, []int{2, 4, 6}) {
		t.Errorf("Filter(even) = %v, want [2 4 6]", got)
	}
}

func TestCounter(t *testing.T) {
	a := NewCounter()
	b := NewCounter()
	for want := 1; want <= 3; want++ {
		if got := a(); got != want {
			t.Errorf("a() = %d, want %d", got, want)
		}
	}
	if got := b(); got != 1 {
		t.Errorf("второй счётчик должен быть независимым: b() = %d, want 1", got)
	}
}

func TestCompose(t *testing.T) {
	inc := func(x int) int { return x + 1 }
	sq := func(x int) int { return x * x }
	f := Compose(sq, inc)
	if f == nil {
		t.Fatal("Compose вернула nil")
	}
	if got := f(3); got != 16 {
		t.Errorf("Compose(sq, inc)(3) = %d, want 16 (сначала inc, потом sq)", got)
	}
}
