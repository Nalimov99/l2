package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	input := []string{
		"абв", "вба", "бав", "БВА",
		"непопал",
		"пятак", "пятка", "пятка", "тяпка",
		"листок", "слиток", "столик",
	}

	fmt.Println(FindAnagrams(input))
}

func FindAnagrams(input []string) *map[string][]string {
	result := make(map[string][]string)
	anagramUnique := make(map[string]map[string]bool)

	for _, item := range input {
		wordSlice := strings.Split(item, "")
		sort.Strings(wordSlice)
		word := strings.ToLower(strings.Join(wordSlice, ""))
		item = strings.ToLower(item)

		if _, ok := anagramUnique[word]; ok {
			anagramUnique[word][item] = false
			continue
		}

		anagramUnique[word] = make(map[string]bool)
		anagramUnique[word][item] = true
	}

	for _, item := range anagramUnique {
		if len(item) <= 1 {
			continue
		}

		var firstAnagram string
		slice := make([]string, 0, len(item))
		for word, value := range item {
			if value {
				firstAnagram = word
			}
			slice = append(slice, word)
		}

		sort.Slice(slice, func(i, j int) bool {
			return slice[i] > slice[j]
		})
		result[firstAnagram] = slice
	}
	return &result
}
