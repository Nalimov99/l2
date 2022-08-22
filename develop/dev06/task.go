package main

import (
	"dev06/mancut"
	"flag"
	"fmt"
	"log"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	f := flag.Int("f", -1, "-f=1")
	d := flag.String("d", "\t", "-d=\" \"")
	s := flag.Bool("s", false, "-s")
	flag.Parse()
	filepath := flag.Arg(0)

	m := mancut.Mancut{
		Path: filepath,
		CutFlags: mancut.CutFlags{
			Delimiter: *d,
			Field:     *f,
			S:         *s,
		},
	}

	res, err := m.Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
