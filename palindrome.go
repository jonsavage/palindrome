package palindrome

func IsPalindrome(candidate string) bool {
    if len(candidate) == 0 {
        return false
    }
    if len(candidate) == 1 {
        return true
    }

    return candidate[0] == candidate[len(candidate)-1]
}
