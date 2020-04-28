package sexpr

import (
	"reflect"
	"testing"
)

// 测试验证对一个复杂数据值进行编码和解码会产生相同的结果

// 该测试不对编码输出做出直接断言，因为该输出取决于映射迭代顺序，该顺序是不确定的
// 可以通过使用 -v 参数运行测试来检查 t.Log 语句的输出：
// go test -v
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// 编码
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal Failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// 解码
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// 检查相等性
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// 格式化
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIndent() = %s\n", data)
}
