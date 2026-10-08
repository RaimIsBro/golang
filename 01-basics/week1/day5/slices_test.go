package day5

import (
	"slices"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := [][2][]int{
		{{}, {}},
		{{1}, {1}},
		{{1, 2}, {2, 1}},
		{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		s := slices.Clone(tt[0])
		Reverse(s)
		if !slices.Equal(s, tt[1]) {
			t.Errorf("Reverse(%v) -> %v, want %v", tt[0], s, tt[1])
		}
	}
}

func TestUnique(t *testing.T) {
	in := []int{3, 1, 3, 2, 1, 4}
	orig := slices.Clone(in)
	want := []int{3, 1, 2, 4}
	if got := Unique(in); !slices.Equal(got, want) {
		t.Errorf("Unique(%v) = %v, want %v", orig, got, want)
	}
	if !slices.Equal(in, orig) {
		t.Errorf("Unique изменила входной слайс: %v, было %v", in, orig)
	}
}

func TestChunk(t *testing.T) {
	got := Chunk([]int{1, 2, 3, 4, 5}, 2)
	want := [][]int{{1, 2}, {3, 4}, {5}}
	if len(got) != len(want) {
		t.Fatalf("Chunk(..., 2) = %v, want %v", got, want)
	}
	for i := range want {
		if !slices.Equal(got[i], want[i]) {
			t.Fatalf("Chunk(..., 2) = %v, want %v", got, want)
		}
	}
	if got := Chunk(nil, 3); len(got) != 0 {
		t.Errorf("Chunk(nil, 3) = %v, want пустой результат", got)
	}
	if got := Chunk([]int{1, 2}, 0); got != nil {
		t.Errorf("Chunk(..., 0) = %v, want nil", got)
	}
}

// Ловушка общего базового массива: append в кусок не должен портить соседний кусок.
func TestChunkIndependent(t *testing.T) {
	chunks := Chunk([]int{1, 2, 3, 4}, 2)
	if len(chunks) != 2 || len(chunks[1]) == 0 {
		t.Fatalf("Chunk([1 2 3 4], 2) = %v, want [[1 2] [3 4]]", chunks)
	}
	_ = append(chunks[0], 99)
	if chunks[1][0] != 3 {
		t.Errorf("append в chunks[0] испортил chunks[1]: %v (подсказка: полный срез s[lo:hi:max] или копия)", chunks)
	}
}

func TestRemoveAt(t *testing.T) {
	in := []int{10, 20, 30, 40}
	orig := slices.Clone(in)
	if got := RemoveAt(in, 1); !slices.Equal(got, []int{10, 30, 40}) {
		t.Errorf("RemoveAt(%v, 1) = %v, want [10 30 40]", orig, got)
	}
	if !slices.Equal(in, orig) {
		t.Errorf("RemoveAt изменила входной слайс: %v, было %v", in, orig)
	}
	if got := RemoveAt(orig, 10); !slices.Equal(got, orig) {
		t.Errorf("RemoveAt с индексом вне диапазона = %v, want копию входа %v", got, orig)
	}
}
