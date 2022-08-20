package main

import (
	"dev03/mansort"
	"flag"
	"fmt"
	"os"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	k := flag.Int("k", -1, "-k=1")
	n := flag.Bool("n", false, "-n=true")
	r := flag.Bool("r", false, "-r=true")
	u := flag.Bool("u", false, "-u=true")
	flag.Parse()
	file := flag.Arg(0)

	sf := mansort.SortFlags{
		K: *k,
		N: *n,
		R: *r,
		U: *u,
	}

	dat, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("open %s: no such file or directory\n", file)
		fmt.Println("usage: go run --flags . filepath")
		os.Exit(0)
	}

	sorter := mansort.NewSort(dat, sf)
	sorter.Sort()
	fmt.Println(sorter.Lines())
}
