package palindrome

func IsPalindrome(candidate string) bool {
    if len(candidate) == 0 {
        return false
    }
    if len(candidate) == 1 {
        return true
    }
    return isMirrored(candidate)
}

func isMirrored(candidate string) bool {
    if len(candidate) <= 1 {
        return true
    }
    if firstAndLastCharactersMatch(candidate) {
        return isMirrored(trimFirstAndLastCharacters(candidate))
    }
    return false
}

func firstAndLastCharactersMatch(candidate string) bool {
    return candidate[0] == candidate[len(candidate)-1]
}

func trimFirstAndLastCharacters(aString string) string {
    return aString[1:len(aString) - 1]
}
