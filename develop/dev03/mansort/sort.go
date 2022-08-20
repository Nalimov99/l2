package mansort

import (
	"strings"
)

// SortFlags содержить в себе значения переданных флагов
type SortFlags struct {
	K       int
	N, R, U bool
}

// Sort содержит в себе зависимости и методы для сортировки
type Sort struct {
	Data        []byte
	sortedLines []string
	SortFlags
}

// NewSort знает как инициализировать Sort
func NewSort(data []byte, sortFlags SortFlags) *Sort {
	return &Sort{
		Data:      data,
		SortFlags: sortFlags,
	}
}

// Sort знает как выбрать нужный метод сортировки
// результат сортировки будет записан в структуру Sort.sortedLines
func (s *Sort) Sort() {
	lines := strings.Split(string(s.Data), "\n")

	if s.N {
		s.sortedLines = numSort(lines, s.K)
		return
	}

	s.sortedLines = defaultSort(lines, s.K)
}

// Lines знает какие методы приминить для получение отсортированных данных.
// Возвращаемым значением является отсортированная строка, и если необходимо
// с reverse/unique модификациями
func (s *Sort) Lines() string {
	s.uniqueLines()
	s.reverseLines()

	return strings.Join(s.sortedLines, "\n")
}

// uniqueLines знает как проверить слайс Sorted.sortedLines на уникальность
func (s *Sort) uniqueLines() {
	if !s.U {
		return
	}

	hashUniq := make(map[string]struct{})
	newLines := make([]string, 0, len(s.sortedLines))

	for _, item := range s.sortedLines {
		if _, ok := hashUniq[item]; ok {
			continue
		}

		hashUniq[item] = struct{}{}
		newLines = append(newLines, item)
	}

	s.sortedLines = newLines
}

// reverseLines знает как перевернуть слайс Sorted.sortedLines
func (s *Sort) reverseLines() {
	if !s.R {
		return
	}

	for i, j := 0, len(s.sortedLines)-1; i < j; i, j = i+1, j-1 {
		s.sortedLines[i], s.sortedLines[j] = s.sortedLines[j], s.sortedLines[i]
	}
}
