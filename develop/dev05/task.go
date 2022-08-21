package main

import (
	"dev05/mangrep"
	"flag"
	"fmt"
	"log"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	A := flag.Int("A", -1, "-A=1")
	B := flag.Int("B", -1, "-B=1")
	C := flag.Int("C", -1, "-C=1")
	c := flag.Bool("c", false, "-c")
	i := flag.Bool("i", false, "-i")
	v := flag.Bool("v", false, "-v")
	F := flag.Bool("F", false, "-F")
	n := flag.Bool("n", false, "-n")
	flag.Parse()
	pattern := flag.Arg(0)
	filepath := flag.Arg(1)
	g := mangrep.Grep{
		Path:    filepath,
		Pattern: pattern,
		GrepFlags: mangrep.GrepFlags{
			A:     *A,
			B:     *B,
			C:     *C,
			Count: *c,
			I:     *i,
			V:     *v,
			F:     *F,
			N:     *n,
		},
	}

	if err := g.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.Result())
}
