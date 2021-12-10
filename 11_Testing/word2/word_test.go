package word

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aba", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
		{"中国中", true},
		{"中过", false},
		{"șȘ", true},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf(`IsPalindrome(%q) = %v`, test.input, got)
		}
	}
}

func randomPalindrome(rng *rand.Rand) string {
	length := rng.Intn(25) // [0, 25)
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		r := rune(rng.Intn(1000)) // \u0000 ~ \u0999
		runes[i] = r
		runes[length-i-1] = r
	}
	return string(runes)
}

func nonRandomPalindrome(rng *rand.Rand) string {
	length := rng.Intn(23) + 2 // [2, 25)
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		r := rune(rng.Intn(1000)) // \u0000 ~ \u0999
		runes[i] = r
		runes[length-i-1] = r
	}
	randOffset := rand.Intn(25) + 1
	randIdx := rand.Intn(length / 2)
	runes[randIdx] = rune('a' + randOffset)
	runes[length-1-randIdx] = rune('A' + randOffset - 1)
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		s := randomPalindrome(rng)
		if !IsPalindrome(s) {
			t.Errorf(`IsPalindrome(%q) = false`, s)
		}
	}
}

func TestNonRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		s := nonRandomPalindrome(rng)
		if IsPalindrome(s) {
			t.Errorf(`IsPalindrome(%q) = true`, s)
		}
	}
}
