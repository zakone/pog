package sexpr

import (
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
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

	data, _ := Marshal(&strangelove)
	r := bytes.NewReader(data)
	var strangelove2 Movie
	err := NewDecoder(r).Decode(&strangelove2)
	if err != nil {
		t.Errorf("Decode error happen: %t", err)
	}
	if strangelove2.Title != strangelove.Title {
		t.Errorf("Decode failed, value %s is different with %s", strangelove2.Title, strangelove.Title)
	}
}
