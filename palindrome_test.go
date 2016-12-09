package palindrome

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_IsPalindrome_WithEmptyString_ReturnsFalse(t *testing.T) {
    result := IsPalindrome("")

    assert.False(t, result, "Empty strings are not palindromes")
}

func Test_IsPalindrome_WithSingleCharacterString_ReturnsTrue(t *testing.T) {
    result := IsPalindrome("a")

    assert.True(t, result, "Single character strings are palindromes")
}
