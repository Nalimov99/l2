package mansort

import (
	"sort"
	"strconv"
	"strings"
)

// defaultSort сортировка без переданных флагов
func defaultSort(lines []string, key int) []string {
	if key < 0 {
		sorted := make([]string, len(lines))
		copy(sorted, lines)
		sort.Strings(sorted)

		return sorted
	}

	sortedStr := make([][]string, 0)
	for _, line := range lines {
		sortedStr = append(sortedStr, strings.Split(line, " "))
	}

	sort.SliceStable(sortedStr, func(i, j int) bool {
		if key > len(sortedStr[i])-1 || key > len(sortedStr[j])-1 {
			return true
		}

		stri := strings.Join(sortedStr[i][key:], " ")
		strj := strings.Join(sortedStr[j][key:], " ")

		return stri < strj
	})

	res := make([]string, 0, len(sortedStr))
	for _, item := range sortedStr {
		res = append(res, strings.Join(item, " "))
	}

	return res
}

// numSort сортирует слайс по числовому значению.
// Если на позиции флага --k не окажеться цифры,
// то данная строка будет идти первой в возвращаемом слайсе.
func numSort(lines []string, key int) []string {
	sortedNums := make([][]string, 0)

	for _, line := range lines {
		sortedNums = append(sortedNums, strings.Split(line, " "))
	}

	var k int
	if key > 0 {
		k = key
	}

	sort.SliceStable(sortedNums, func(i, j int) bool {
		if k > len(sortedNums[i])-1 || k > len(sortedNums[j])-1 {
			return true
		}

		wordi := sortedNums[i][k]
		wordj := sortedNums[j][k]

		numi, erri := strconv.Atoi(wordi)
		numj, errj := strconv.Atoi(wordj)
		if erri != nil || errj != nil {
			return true
		}

		return numj < numi
	})

	res := make([]string, 0, len(sortedNums))
	for _, item := range sortedNums {
		res = append(res, strings.Join(item, " "))
	}

	return res
}
