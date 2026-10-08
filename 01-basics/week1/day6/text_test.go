package day6

import (
	"maps"
	"slices"
	"testing"
)

func TestRuneLen(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"", 0},
		{"Go", 2},
		{"Привет", 6},
		{"Go — это 🐹", 10},
	}
	for _, tt := range tests {
		if got := RuneLen(tt.in); got != tt.want {
			t.Errorf("RuneLen(%q) = %d, want %d (а len() = %d байт)", tt.in, got, tt.want, len(tt.in))
		}
	}
}

func TestReverseString(t *testing.T) {
	tests := map[string]string{
		"":       "",
		"abc":    "cba",
		"Привет": "тевирП",
		"Go🐹":    "🐹oG",
	}
	for in, want := range tests {
		if got := ReverseString(in); got != want {
			t.Errorf("ReverseString(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	yes := []string{"", "шалаш", "А роза упала на лапу Азора", "Was it a car or a cat I saw?"}
	no := []string{"Go", "Привет, мир"}
	for _, s := range yes {
		if !IsPalindrome(s) {
			t.Errorf("IsPalindrome(%q) = false, want true", s)
		}
	}
	for _, s := range no {
		if IsPalindrome(s) {
			t.Errorf("IsPalindrome(%q) = true, want false", s)
		}
	}
}

func TestWordCount(t *testing.T) {
	text := "Мама мыла раму. Раму мыла мама! go, Go, GO?"
	want := map[string]int{"мама": 2, "мыла": 2, "раму": 2, "go": 3}
	if got := WordCount(text); !maps.Equal(got, want) {
		t.Errorf("WordCount(%q) = %v, want %v", text, got, want)
	}
	if got := WordCount("   "); len(got) != 0 {
		t.Errorf("WordCount для пустого текста = %v, want пустую map", got)
	}
}

func TestTopWords(t *testing.T) {
	counts := map[string]int{"кот": 3, "пёс": 5, "енот": 3, "мышь": 1}
	tests := []struct {
		n    int
		want []string
	}{
		{1, []string{"пёс"}},
		{3, []string{"пёс", "енот", "кот"}},
		{10, []string{"пёс", "енот", "кот", "мышь"}},
		{0, []string{}},
	}
	for _, tt := range tests {
		if got := TopWords(counts, tt.n); !slices.Equal(got, tt.want) {
			t.Errorf("TopWords(n=%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}
