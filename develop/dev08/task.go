package main

import (
	"dev08/shell"
	"os"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
*/

func main() {
	shell.StartShell(os.Stdin, os.Stdout)
}
