package main

import (
	"dev09/wget"
	"flag"
	"fmt"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	r := flag.Bool("r", false, "-r")
	flag.Parse()

	if _, err := wget.NewWget(flag.Arg(0), wget.Flags{R: *r}); err != nil {
		fmt.Println(err)
	}
}
