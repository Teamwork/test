package test

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestErrorContains(t *testing.T) {
	cases := []struct {
		err      error
		str      string
		expected bool
	}{
		{errors.New("Hello"), "Hello", true},
		{errors.New("Hello, world"), "world", true},
		{nil, "", true},

		{errors.New("Hello, world"), "", false},
		{errors.New("Hello, world"), "mars", false},
		{nil, "hello", false},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.err), func(t *testing.T) {
			out := ErrorContains(tc.err, tc.str)
			if out != tc.expected {
				t.Errorf("\nout:      %#v\nexpected: %#v\n", out, tc.expected)
			}
		})
	}
}

func TestTempFile(t *testing.T) {
	f, clean := TempFile(t, "hello\nworld")

	_, err := os.Stat(f)
	if err != nil {
		t.Fatal(err)
	}

	clean()

	_, err = os.Stat(f)
	if err == nil {
		t.Fatal(err)
	}
}

func TestNormalizeIndent(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{
			"\t\twoot\n\t\twoot\n",
			"woot\nwoot",
		},
		{
			"\t\twoot\n\t\t woot",
			"woot\n woot",
		},
		{
			"\t\twoot\n\t\t\twoot",
			"woot\n\twoot",
		},
		{
			"woot\n\twoot",
			"woot\n\twoot",
		},
		{
			"  woot\n\twoot",
			"woot\n\twoot",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			out := NormalizeIndent(tc.in)
			if out != tc.want {
				t.Errorf("\nout:  %#v\nwant: %#v\n", out, tc.want)
			}
		})
	}
}
