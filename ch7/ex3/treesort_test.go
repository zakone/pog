package treesort

import "testing"
import "math/rand"
import "fmt"

func TestTreeSort(t *testing.T) {
	data := make([]int, 10)
	for i := range data {
		data[i] = rand.Int() % 10
	}
	tree := Sort(data)
	if tree.String() != fmt.Sprintf("%v", data) {
		t.Errorf("string output incorrect¥n")
		t.Errorf("%s", tree.String())
		t.Errorf("%s", fmt.Sprintf("%v", data))
	}
}
