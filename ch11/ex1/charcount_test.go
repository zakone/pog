package charcount

import "testing"

func TestCharcount(t *testing.T) {
	count, _ := Charcount("あああいいう")
	if count['あ'] != 3 || count['い'] != 2 || count['う'] != 1 {
		t.Errorf("count wrong number: 'あ': %d, 'い': %d, 'う': %d", count['あ'], count['い'], count['う'])
	}
}

func TestCharcountVaild(t *testing.T) {
	_, invalid := Charcount("あいうえお\x80")
	if invalid != 1 {
		t.Errorf("count invalid number: %d, should be 1", invalid)
	}
}

func TestCharcountEmpty(t *testing.T) {
	count, _ := Charcount("")
	if len(count) > 0 {
		t.Errorf("count empty string: %d, should be 0", len(count))
	}
}
