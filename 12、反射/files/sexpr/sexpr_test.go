package sexpr

import "testing"

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

	// 格式化
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIndent() = %s\n", data)
}
