package popcount_test

import (
	"testing"
	"./popcount"
	"sync"
)


func TestPopCount(t *testing.T) {
	
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			count := popcount.PopCount(0x1234567890ABCDEF)
			if got, want := count, 32; got != want {
				t.Errorf("Popcount Error: %d, want %d", got, want)
			}
		}()
	}
	wg.Wait()

}