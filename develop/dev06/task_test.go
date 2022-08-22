package main

import (
	"dev06/mancut"
	"testing"
)

func TestMancut(t *testing.T) {
	mt := mancutTest{
		mc: &mancut.Mancut{
			Path: "./assets/test.txt",
		},
	}

	t.Log("RUN MANCUT TESTS")
	t.Run("CUT -F=1", mt.TestFlagF1)
	t.Run("CUT -F=4", mt.TestFlagF4)
	t.Run("CUT -F=4 -D=' '", mt.TestSpaceDelimiterF4)
	t.Run("CUT -F=1 -D-'-'", mt.TestDashDelimiterF1)
	t.Run("CUT -D=' ' -S -F=1", mt.TestSpaceDelimeterSF1)
	t.Run("CUT -D='-' -S -F=5", mt.TestDashDelimeterSF5)
}

func checkCutResult(t *testing.T, want, got string, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("want:\n%s\n\ngot:\n%s", want, got)
	}
}

type mancutTest struct {
	mc *mancut.Mancut
}

func (m *mancutTest) TestFlagF1(t *testing.T) {
	m.mc.CutFlags = mancut.CutFlags{
		Field: 1,
	}

	want := `asdbcacda
fff ffa ffb
a
a-c-b-c
a\nb`

	got, err := m.mc.Result()
	checkCutResult(t, want, got, err)
}

func (m *mancutTest) TestFlagF4(t *testing.T) {
	m.mc.CutFlags = mancut.CutFlags{
		Field: 4,
	}

	want := `asdbcacda
fff ffa ffb

a-c-b-c
a\nb`

	got, err := m.mc.Result()
	checkCutResult(t, want, got, err)
}

func (m *mancutTest) TestSpaceDelimiterF4(t *testing.T) {
	m.mc.CutFlags = mancut.CutFlags{
		Field:     4,
		Delimiter: " ",
	}

	want := `asdbcacda

a	b
a-c-b-c
a\nb`

	got, err := m.mc.Result()
	checkCutResult(t, want, got, err)
}

func (m *mancutTest) TestDashDelimiterF1(t *testing.T) {
	m.mc.CutFlags = mancut.CutFlags{
		Field:     1,
		Delimiter: "-",
	}

	want := `asdbcacda
fff ffa ffb
a	b
a
a\nb`

	got, err := m.mc.Result()
	checkCutResult(t, want, got, err)
}

func (m *mancutTest) TestSpaceDelimeterSF1(t *testing.T) {
	m.mc.CutFlags = mancut.CutFlags{
		Delimiter: " ",
		Field:     1,
		S:         true,
	}

	want := "fff"
	got, err := m.mc.Result()
	checkCutResult(t, want, got, err)
}

func (m *mancutTest) TestDashDelimeterSF5(t *testing.T) {
	m.mc.CutFlags = mancut.CutFlags{
		Delimiter: "-",
		Field:     5,
		S:         true,
	}

	want := ""
	got, err := m.mc.Result()
	checkCutResult(t, want, got, err)
}
