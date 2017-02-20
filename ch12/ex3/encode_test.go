package encode

import (
	"testing"
)

func TestBooleanFlaseEncode(t *testing.T) {
	res, _ := Marshal(false)
	if string(res) != "nil" {
		t.Errorf("marshal false boolean should be 'nil', but %s", res)
	}

}

func TestBooleanTrueEncode(t *testing.T) {
	res, _ := Marshal(true)
	if string(res) != "t" {
		t.Errorf("marshal false boolean should be 't', but %s", res)
	}
}

func TestComplexEncode(t *testing.T) {
	res, _ := Marshal(complex(1.0, 2.0))
	if string(res) != "#C(1.0 2.0)" {
		t.Errorf("marshal false boolean should be '#C(1.0 2.0)', but %s", res)
	}
}

func TestInterfaceEncode(t *testing.T) {
	var i interface{} = 3
	res, _ := Marshal(i)
	if string(res) != "3" {
		t.Errorf("marshal false boolean should be '3', but %s", res)
	}
}