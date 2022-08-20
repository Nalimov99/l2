package main

import "testing"

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		arg, want string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"", ""},
		{`qwe\4\5`, "qwe45"},
		{`qwe\45`, "qwe44444"},
		{`qwe\\5`, `qwe\\\\\`},
		{"a10b", "aaaaaaaaaab"},
		{`a\b\s\d\\3`, `a\b\s\d\\\`},
	}

	for _, item := range testCases {
		str, err := UnpackString(item.arg)
		if err != nil {
			t.Fatalf("err should be nil. arg: %s; got: %s; want: %s; err: %v",
				item.arg,
				str,
				item.want,
				err,
			)
		}
	}

	_, err := UnpackString("45")
	if err != ErrInvalidString {
		t.Fatalf("unexpected error: %v", err)
	}
}
