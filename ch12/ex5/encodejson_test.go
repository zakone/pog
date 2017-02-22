package encodejson

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJsonEncode(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		Param1          interface{}
		Param2          *int
	}
	var y interface{} = 3
	//!-movie
	//!+strangelove
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
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
		Param1: y,
		Param2: nil,
	}
	b, err := Marshal(strangelove)
	if err != nil {
		t.Errorf("marshal json failed with err: %t", err)
	}
	var strangelove2 Movie
	if err = json.Unmarshal(b, &strangelove2); err != nil {
		t.Errorf("unmarshal json failed with err: %t", err)
	}
	if strangelove2.Title != strangelove.Title {
		t.Errorf("json encode Movie.Title: %s diffrent with orignal Title: %s", strangelove2.Title, strangelove.Title)
	}
	if strangelove2.Subtitle != strangelove.Subtitle {
		t.Errorf("json encode Movie.Subtitle: %s diffrent with orignal Subtitle: %s", strangelove2.Subtitle, strangelove.Subtitle)
	}
	if strangelove2.Year != strangelove.Year {
		t.Errorf("json encode Movie.Year: %s diffrent with orignal Year: %s", strangelove2.Year, strangelove.Year)
	}
	if strangelove2.Color != strangelove.Color {
		t.Errorf("json encode Movie.Color: %s diffrent with orignal Color: %s", strangelove2.Color, strangelove.Color)
	}
	if !reflect.DeepEqual(strangelove.Oscars, strangelove2.Oscars) {
		t.Errorf("json encode Movie.Oscars: %s diffrent with orignal Oscars: %s", strangelove2.Oscars, strangelove.Oscars)
	}

}
