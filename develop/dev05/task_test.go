package main

import (
	"dev05/mangrep"
	"path"
	"testing"
)

func TestGrep(t *testing.T) {
	tg := testGrep{
		grep: &mangrep.Grep{
			Path: path.Join(".", "assets", "test.txt"),
		},
	}

	t.Log("RUN GREP TESTS")
	t.Run("GREP WITHOUT FLAGS", tg.TestGrep)
	t.Run("GREP B=1", tg.TestGrepB1)
	t.Run("GREP A=1", tg.TestGrepA1)
	t.Run("GREP A=1 B=1", tg.TestGrepA1B1)
	t.Run("GREP C=1", tg.TestGrepC1)
	t.Run("GREP C=1 B=1 A=1", tg.TestGrepC1A1B1)
	t.Run("GREP COUNT", tg.TestGrepCount)
	t.Run("GREP V", tg.TestGrepInvert)
	t.Run("GREP F", tg.TestGrepFixed)
	t.Run("GREP N", tg.TestGrepNumLines)
}

type testGrep struct {
	grep *mangrep.Grep
}

func gotWantLog(t *testing.T, want, got string) {
	t.Helper()

	if got != want {
		t.Fatalf("want:\n%s\n\ngot:\n%s", want, got)
	}
}

func (tg *testGrep) TestGrep(t *testing.T) {
	tg.grep.Pattern = "фыв"
	tg.grep.Start()
	want := `-A - "after" печатать +N строк после совпадения фыв
-i - "ignore-case" (игнорировать регистр) фыв`

	gotWantLog(t, want, tg.grep.Result())
}

func (tg *testGrep) TestGrepB1(t *testing.T) {
	tg.grep.GrepFlags = mangrep.GrepFlags{
		B: 1,
	}
	tg.grep.Start()
	want := `Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения фыв
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр) фыв`

	gotWantLog(t, want, tg.grep.Result())
}

func (tg *testGrep) TestGrepA1(t *testing.T) {
	tg.grep.GrepFlags = mangrep.GrepFlags{
		A: 1,
	}
	tg.grep.Start()
	want := `-A - "after" печатать +N строк после совпадения фыв
-B - "before" печатать +N строк до совпадения
-i - "ignore-case" (игнорировать регистр) фыв
-v - "invert" (вместо совпадения, исключать)`

	gotWantLog(t, want, tg.grep.Result())
}

func (tg *testGrep) TestGrepA1B1(t *testing.T) {
	tg.grep.GrepFlags = mangrep.GrepFlags{
		A: 1,
		B: 1,
	}
	tg.grep.Start()
	want := `Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения фыв
-B - "before" печатать +N строк до совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр) фыв
-v - "invert" (вместо совпадения, исключать)`

	gotWantLog(t, want, tg.grep.Result())
}

func (tg *testGrep) TestGrepC1(t *testing.T) {
	tg.grep.GrepFlags = mangrep.GrepFlags{
		C: 1,
	}
	tg.grep.Start()
	want := `Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения фыв
-B - "before" печатать +N строк до совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр) фыв
-v - "invert" (вместо совпадения, исключать)`

	gotWantLog(t, want, tg.grep.Result())
}

func (tg *testGrep) TestGrepC1A1B1(t *testing.T) {
	tg.grep.GrepFlags = mangrep.GrepFlags{
		C: 1,
		A: 1,
		B: 1,
	}
	tg.grep.Start()
	want := `Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения фыв
-B - "before" печатать +N строк до совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр) фыв
-v - "invert" (вместо совпадения, исключать)`

	gotWantLog(t, want, tg.grep.Result())
}

func (tg *testGrep) TestGrepCount(t *testing.T) {
	tg.grep.GrepFlags = mangrep.GrepFlags{
		Count: true,
	}
	tg.grep.Start()
	want := "2"

	gotWantLog(t, want, tg.grep.Result())
}

func (tg *testGrep) TestGrepInvert(t *testing.T) {
	tg.grep.GrepFlags = mangrep.GrepFlags{
		V: true,
	}
	tg.grep.Start()
	want := `Утилита grep

Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).\n

Реализовать поддержку утилитой следующих ключей:
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки`

	gotWantLog(t, want, tg.grep.Result())
}

func (tg *testGrep) TestGrepFixed(t *testing.T) {
	tg.grep.GrepFlags = mangrep.GrepFlags{
		F: true,
	}
	tg.grep.Pattern = `\n`
	tg.grep.Start()
	want := `Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).\n`

	gotWantLog(t, want, tg.grep.Result())
}

func (tg *testGrep) TestGrepNumLines(t *testing.T) {
	tg.grep.GrepFlags = mangrep.GrepFlags{
		N: true,
	}
	tg.grep.Pattern = `фыв`
	tg.grep.Start()
	want := `6. -A - "after" печатать +N строк после совпадения фыв
10. -i - "ignore-case" (игнорировать регистр) фыв`

	gotWantLog(t, want, tg.grep.Result())
}
