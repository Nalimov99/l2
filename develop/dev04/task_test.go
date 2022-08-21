package main

import "testing"

func TestFindAnagrams(t *testing.T) {
	testCase := []string{
		"абв", "вба", "бав", "БВА",
		"непопал",
		"пятак", "пятка", "пятка", "тяпка",
		"листок", "слиток", "столик",
	}

	want := map[string][]string{
		"абв":    {"вба", "бва", "бав", "абв"},
		"пятак":  {"тяпка", "пятка", "пятак"},
		"листок": {"столик", "слиток", "листок"},
	}
	got := *FindAnagrams(testCase)

	for key, slice := range want {
		for i, word := range slice {
			if word != got[key][i] {
				t.Logf("\nwant: %s\ngot:%s", word, got[key][i])
				t.Fatalf("\nwant:%v\ngot:%v", want, got)
			}
		}
	}
}
