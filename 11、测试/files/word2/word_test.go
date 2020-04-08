package word

import (
	"math/rand"
	"testing"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell, lewd did I live", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste", true},
		{"palindrome", false},
		{"desserts", false},
		{"马歇尔歇马", true},
		{"华莱士来华", false},
	}
	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// randomPalindrome 返回一个回文，其长度和内容均来自伪随机数生成器 rng
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // 随机长度不超过 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // 随机 rune 的最大值 \u0999
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	// 初始化伪随机数生成器
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		// fmt.Println(p)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
