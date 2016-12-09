package palindrome

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_IsPalindrome_WithEmptyString_ReturnsFalse(t *testing.T) {
    Assert_IsNotPalindrome(t, "")
}

func Test_IsPalindrome_WithSingleCharacterString_ReturnsTrue(t *testing.T) {
    Assert_IsPalindrome(t, "a")
}

func Test_IsPalindrome_WithTwoCharacterValidPalindrome_ReturnsTrue(t *testing.T) {
    Assert_IsPalindrome(t, "a")
}

func Test_IsPalindrome_WithTwoCharacterInvalidPalindrome_ReturnsFalse(t *testing.T) {
    Assert_IsNotPalindrome(t, "ab")
}

func Test_IsPalindrome_WithThreeCharacterInvalidPalindrome_ReturnsFalse(t *testing.T) {
    Assert_IsNotPalindrome(t, "abc")
}

func Test_IsPalindrome_WithThreeCharacterValidPalindrome_ReturnsTrue(t *testing.T) {
    Assert_IsPalindrome(t, "aba")
}

func Assert_IsPalindrome(t *testing.T, candidate string) {
    result := IsPalindrome(candidate)

    assert.True(t, result, candidate + " is a palindrome")
}

func Assert_IsNotPalindrome(t *testing.T, candidate string) {
    result := IsPalindrome(candidate)

    assert.False(t, result, candidate + " is NOT a palindrome")
}
