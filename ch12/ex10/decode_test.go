package sexpr

import (
	"testing"
)

func TestDecode(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		Published       bool
		Rating          float64
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
		// Published: true,
		Rating: 88.7,
	}

	data, _ := Marshal(strangelove)
	var strangelove2 Movie
	err := Unmarshal(data, strangelove2)
	if err != nil {
		t.Errorf("Decode error happen: %t", err)
	}
	if strangelove2.Published != strangelove.Published {
		t.Errorf("Decode failed, value %t is different with %t", strangelove2.Published, strangelove.Published)
	}
	if strangelove2.Rating != strangelove.Rating {
		t.Errorf("Decode failed, value %t is different with %t", strangelove2.Rating, strangelove.Rating)
	}
}
