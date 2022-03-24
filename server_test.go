package main

import "testing"

func TestIncrementAddr(t *testing.T) {

	//table driven test
	var tests = []struct {
		a string
		b string
	}{
		{":8080", ":8081"},
		{":8991", ":8992"},
	}

	for _, test := range tests {
		want := test.b
		if got := incrementAddr(test.a); got != want {
			t.Errorf("incrementAddr() = %s, want %s", got, want)
		}
	}
}
