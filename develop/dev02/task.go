package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	ErrInvalidString = errors.New("invalid string")
)

func main() {
	fmt.Println(UnpackString(`a4bc2d5e`))
}

func UnpackString(input string) (string, error) {
	validRe, _ := regexp.Compile(`^\d+$`)
	if validRe.Match([]byte(input)) {
		return "", ErrInvalidString
	}

	re, _ := regexp.Compile(`([a-zA-Z]|\\.)(\d+)`)
	groups := re.FindAllStringSubmatch(input, -1)
	str := input
	for _, group := range groups {
		count, _ := strconv.Atoi(group[2])
		repeatSymbol := strings.Replace(group[1], "\\", "", 1)
		str = strings.ReplaceAll(str, group[0], strings.Repeat(repeatSymbol, count))
	}

	reEscapeDigit, _ := regexp.Compile(`(\\)(\d)`)
	str = reEscapeDigit.ReplaceAllString(str, "$2")

	return str, nil
}
