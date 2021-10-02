// Запуск тестов: go test -v
package main

import (
	"fmt"
	"testing"
)

func TestGetWordCount(t *testing.T) {
	got := getWordCount("one two three")
	want := 3
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func TestGetWordCountEmpty(t *testing.T) {
	got := getWordCount("")
	want := 0
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

// Пример табличного теста (table-driven)
func TestGetWordCountTD(t *testing.T) {
	var tests = []struct {
		// name string Имя теста можно задавать явно или генерировать - как показано ниже
		text string
		want int
	}{
		{"word", 1},
		{"", 0},
		{"many many words and symbol -", 6},
	}

	for _, test := range tests {
		name := fmt.Sprintf("case(%s)", test.text)
		t.Run(name, func(t *testing.T) {
			got := getWordCount(test.text)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}
