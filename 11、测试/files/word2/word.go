package word

import "unicode"

// IsPalindrome 回文检测
func IsPalindrome(s string) bool {
	// *+ 11.2 节 2
	// var letters []rune
	// for _, r := range s {
	// 	if unicode.IsLetter(r) {
	// 		letters = append(letters, unicode.ToLower(r))
	// 	}
	// }
	// *- 11.2 节 2

	// *+ 11.2 节 1
	// for i := range letters {
	// 	if letters[i] != letters[len(letters)-1-i] {
	// 		return false
	// 	}
	// }
	// *- 11.2 节 1

	// !+ 11.4 节优化 2
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	// !+ 11.4 节优化 2

	// !+ 11.4 节优化 1
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	// !- 11.4 节优化 1

	return true
}
