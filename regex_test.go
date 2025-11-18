package main

import "testing"

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isMatch("palindromes", "pal*in*d*.")
	}
}

func BenchmarkIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isMatchIterative("palindrome", "pal*in*d*.")
	}
}

func BenchmarkLargeInput(b *testing.B) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	p := "a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*a*"

	for i := 0; i < b.N; i++ {
		isMatchIterative(s, p)
	}
}

