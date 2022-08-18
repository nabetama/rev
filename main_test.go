package main

import "testing"

func TestRev(t *testing.T) {
	var tests = []struct {
		given  string
		expect string
	}{
		{"hoge", "egoh"},
		{"", ""},
		{"hoge\nfuga", "aguf\negoh"},
		{"hoge\tfuga", "aguf\tegoh"},
	}

	for _, tc := range tests {
		t.Run("文字列の反転が出来る", func(t *testing.T) {
			if actual := rev(tc.given); actual != tc.expect {
				t.Errorf("expect %s, but %s", tc.expect, actual)
			}
		})
	}
}
