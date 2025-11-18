// Regular Expression Matcher
// Supports . (matches any single character) and * (matches zero or more of the preceding element)
// It uses recursion with memoization (top-down DP) so it's efficient for typical inputs.
package main

import (
	"fmt"
)

// isMatch returns true if string s matches pattern p.
// Supported special characters:
// '.' matches any single character.
// '*' matches zero or more of the preceding element.
// RECURSIVE + MEMOIZED VERSION
func isMatch(s, p string) bool {
	// memoization map: key is pair (i,j) -> value bool
	type pair struct{ i, j int }
	memo := make(map[pair]bool)

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// If we have computed this state before, return cached result
		key := pair{i, j}
		if v, ok := memo[key]; ok {
			return v
		}

		var ans bool

		// If we have reached the end of the pattern, match only if string also finished
		if j == len(p) {
			ans = (i == len(s))
		} else {
			// Check if current characters match (taking '.' into account)
			firstMatch := i < len(s) && (p[j] == s[i] || p[j] == '.')

			// If there's a '*' as the next pattern character, we have two choices:
			// 1) Treat "x*" as repeating 0 times -> skip "x*"
			// 2) If firstMatch, consume one char from s and keep pattern at j (use '*' again)
			if j+1 < len(p) && p[j+1] == '*' {
				// zero occurrences OR one/more occurrences (if firstMatch)
				ans = dp(i, j+2) || (firstMatch && dp(i+1, j))
			} else {
				// No '*', must consume one character from both if they match
				ans = firstMatch && dp(i+1, j+1)
			}
		}

		memo[key] = ans
		return ans
	}

	return dp(0, 0)
}

// ITERATIVE BOTTOM-UP DP VERSION
// isMatchIterative implements regex matching using an iterative DP table.
func isMatchIterative(s, p string) bool {
	m, n := len(s), len(p)

	// dp[i][j] means s[:i] matches p[:j]
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// empty string matches empty pattern
	dp[0][0] = true
	// support patterns like a*, a*b*, a*b*c*
	for j := 1; j < n; j++ {
		if p[j] == '*' && dp[0][j-1] {
			dp[0][j+1] = true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]

			} else if p[j] == '*' {
				// zero occurrences
				if dp[i+1][j-1] {
					dp[i+1][j+1] = true
				} else if p[j-1] == '.' || p[j-1] == s[i] {
					// one or more occurrences
					if dp[i][j+1] {
						dp[i+1][j+1] = true
					}
				}
			}
		}
	}

	return dp[m][n]
}

// FULL TEST SUITE AND MAIN (RUN TEST)
func main() {
	tests := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"palindromes", "pa*lin*d*.", false},
		{"", ".*", true},
		{"", "", true},
		{"abc", "abc", true},
		{"aaa", "a*a", true},
		{"ab", ".*c", false},
		{"abbb", "ab*", true},
		{"abcd", "d*", false},
		{"aaa", "ab*a*c*a", true},
		{"abc", ".*abc", true},
		{"abc", ".*abcd", false},
	}
	fmt.Println("=== Recursive + Memoized Version ===")
	for _, tc := range tests {
		got := isMatch(tc.s, tc.p)
		fmt.Printf("isMatch(%q, %q) = %v (expected %v)\n", tc.s, tc.p, got, tc.want)
	}
	fmt.Println("\n=== Iterative Bottom-Up DP Version ===")
	for _, tc := range tests {
		got := isMatchIterative(tc.s, tc.p)
		fmt.Printf("isMatchIterative(%q, %q) = %v (expected %v)\n", tc.s, tc.p, got, tc.want)
	}
}

