package wordscounter

import "bufio"
import "bytes"

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(input []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

func (c *LineCounter) Write(input []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}
