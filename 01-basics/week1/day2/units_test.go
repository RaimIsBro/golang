package day2

import (
	"math"
	"testing"
)

func TestWeekdayValues(t *testing.T) {
	if Monday != 0 || Sunday != 6 {
		t.Fatalf("Monday = %d, Sunday = %d; want 0 и 6 (используйте iota)", Monday, Sunday)
	}
}

func TestWeekdayString(t *testing.T) {
	tests := []struct {
		day  Weekday
		want string
	}{
		{Monday, "понедельник"},
		{Wednesday, "среда"},
		{Sunday, "воскресенье"},
		{Weekday(7), "Weekday(7)"},
		{Weekday(-1), "Weekday(-1)"},
	}
	for _, tt := range tests {
		if got := tt.day.String(); got != tt.want {
			t.Errorf("Weekday(%d).String() = %q, want %q", int(tt.day), got, tt.want)
		}
	}
}

func TestIsWeekend(t *testing.T) {
	for d := Monday; d <= Sunday; d++ {
		want := d == Saturday || d == Sunday
		if got := d.IsWeekend(); got != want {
			t.Errorf("%v.IsWeekend() = %v, want %v", d, got, want)
		}
	}
}

func TestByteSizeConstants(t *testing.T) {
	if KB != 1024 || MB != 1024*1024 || GB != 1024*1024*1024 {
		t.Errorf("KB=%d MB=%d GB=%d; ожидаются степени 1024 (подсказка: iota и сдвиг <<)", KB, MB, GB)
	}
}

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		in   int64
		want string
	}{
		{0, "0 B"},
		{512, "512 B"},
		{1023, "1023 B"},
		{1024, "1.0 KB"},
		{1536, "1.5 KB"},
		{10 * 1024 * 1024, "10.0 MB"},
		{3*1024*1024*1024 + 512*1024*1024, "3.5 GB"},
	}
	for _, tt := range tests {
		if got := FormatBytes(tt.in); got != tt.want {
			t.Errorf("FormatBytes(%d) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct{ c, f float64 }{
		{0, 32},
		{100, 212},
		{-40, -40},
		{36.6, 97.88},
	}
	for _, tt := range tests {
		if got := CelsiusToFahrenheit(tt.c); math.Abs(got-tt.f) > 1e-9 {
			t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tt.c, got, tt.f)
		}
	}
}
