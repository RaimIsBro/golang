package day1

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "латиница", in: "Gopher", want: "Привет, Gopher!"},
		{name: "кириллица", in: "Раим", want: "Привет, Раим!"},
		{name: "пустое имя", in: "", want: "Привет, мир!"},
		{name: "только пробелы", in: "   ", want: "Привет, мир!"},
		{name: "пробелы по краям", in: "  Аня ", want: "Привет, Аня!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.in); got != tt.want {
				t.Errorf("Greet(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
