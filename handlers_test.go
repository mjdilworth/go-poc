package main

import "testing"

func TestVerifyUserPass(t *testing.T) {

	//table driven test
	var tests = []struct {
		a string
		b string
		c bool
	}{
		{"user1", "pass1", true},
		{"user2", "pass2", true},
	}

	for _, test := range tests {
		want := test.c
		if got := verifyUserPass(test.a, test.b); got != want {
			t.Errorf("verifyUserPass() = %t, want %t", got, want)
		}

	}
}
