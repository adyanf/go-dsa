package knowing_what_to_track

func PermutePalindrome(st string) bool {
	charTracker := make(map[rune]int)

	for _, char := range st {
		if _, ok := charTracker[char]; ok {
			charTracker[char]++
		} else {
			charTracker[char] = 1
		}
	}

	oddCount := 0
	for _, count := range charTracker {
		if count%2 != 0 {
			oddCount++
		}
	}
	return oddCount <= 1
}

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charTracker := make(map[rune]int)

	for _, char := range s1 {
		if _, ok := charTracker[char]; ok {
			charTracker[char]++
		} else {
			charTracker[char] = 1
		}
	}

	for _, char := range s2 {
		if _, ok := charTracker[char]; ok {
			charTracker[char]--

			if charTracker[char] == 0 {
				delete(charTracker, char)
			}
		} else {
			return false
		}
	}

	return true
}
