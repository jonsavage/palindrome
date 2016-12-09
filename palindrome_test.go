package palindrome

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_IsPalindrome_WithEmptyString_ReturnsFalse(t *testing.T) {
    result := IsPalindrome("")

    assert.False(t, result, "Empty strings are not palindromes")
}
