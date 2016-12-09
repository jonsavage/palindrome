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

func Test_IsPalindrome_WithTwoCharacterValidPalindrome_ReturnsTrue(t *testing.T) {
    result := IsPalindrome("aa")

    assert.True(t, result, "'aa' is a palindrome")
}

func Test_IsPalindrome_WithTwoCharacterInvalidPalindrome_ReturnsFalse(t *testing.T) {
    result := IsPalindrome("ab")

    assert.False(t, result, "'ab' is not a palindrome")
}

func Test_IsPalindrome_WithThreeCharacterInvalidPalindrome_ReturnsFalse(t *testing.T) {
    result := IsPalindrome("abc")

    assert.False(t, result, "'abc' is not a palindrome") 
}

func Test_IsPalindrome_WithThreeCharacterValidPalindrome_ReturnsTrue(t *testing.T) {
    result := IsPalindrome("aba")

    assert.True(t, result, "'aba' is a palindrome") 
}
