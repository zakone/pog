package encodenonzero

import (
	"testing"
)

func TestEncodeNonZeroInt(t *testing.T) {
	v1 := 3
	b, err := Marshal(v1)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if string(b) != "3" {
		t.Errorf("encode non zero value %s, should be 3", b)
	}
	v2 := 0
	b, err = Marshal(v2)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if len(b) != 0 {
		t.Errorf("encode non zero value %s, should be nil", b)
	}
}

func TestEncodeNonZeroString(t *testing.T) {
	v1 := "aaa"
	b, err := Marshal(v1)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if string(b) != "\"aaa\"" {
		t.Errorf("encode non zero value %s, should be aaa", b)
	}
	v2 := ""
	b, err = Marshal(v2)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if len(b) != 0 {
		t.Errorf("encode non zero value %s, should be nil", b)
	}
}

func TestEncodeNonZeroPtr(t *testing.T) {
	v1 := 1
	p1 := &v1
	b, err := Marshal(p1)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if string(b) != "1" {
		t.Errorf("encode non zero value %s, should be 1", b)
	}
	var p2 *int
	p2 = nil
	b, err = Marshal(p2)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if len(b) != 0 {
		t.Errorf("encode non zero value %s, should be nil", b)
	}
}

func TestEncodeNonZeroArray(t *testing.T) {
	var v1 = []int{1, 2, 3}
	b, err := Marshal(v1)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if string(b) != "(1 2 3)" {
		t.Errorf("encode non zero value %s, should be (1 2 3)", b)
	}
	var v2 []int
	b, err = Marshal(v2)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if len(b) != 0 {
		t.Errorf("encode non zero value %s, should be nil", b)
	}
}

func TestEncodeNonZeroMap(t *testing.T) {
	var v1 = map[string]int{
		"key1": 1,
		"key2": 2,
	}
	b, err := Marshal(v1)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if string(b) != "((\"key1\" 1) (\"key2\" 2))" && string(b) != "(\"key2\" 2)) ((\"key1\" 1)" {
		t.Errorf("encode non zero value %s, should be ((key1 1) (key2 2))", b)
	}
	var v2 map[string]int
	b, err = Marshal(v2)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if len(b) != 0 {
		t.Errorf("encode non zero value %s, should be nil", b)
	}
}

func TestEncodeNonZeroStruct(t *testing.T) {
	type S1 struct {
		P1 int
		P2 bool
		P3 float64
	}
	s1 := S1{
		P1: 1,
		P2: true,
		P3: 1.5,
	}
	b, err := Marshal(s1)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if string(b) != "((P1 1) (P2 t) (P3 1.5))" {
		t.Errorf("encode non zero value %s, should be ((P1 1) (P2 t) (P3 1.5))", b)
	}
	type S2 struct {
		P1 int
		P2 bool
		P3 float64
	}
	s2 := S2{
		P1: 0,
		P2: false,
		P3: 0.0,
	}
	b, err = Marshal(s2)
	if err != nil {
		t.Errorf("encode error: %t", err)
	}
	if string(b) != "()" {
		t.Errorf("encode non zero value %s, should be ()", b)
	}
}
