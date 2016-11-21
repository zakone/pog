package writercounter

import "io"

type WriterCounter struct {
	count  int64
	writer io.Writer
}

func (counter *WriterCounter) Write(buf []byte) (int, error) {
	n, err := counter.writer.Write(buf)
	counter.count = int64(n)
	return n, err
}

func (counter *WriterCounter) Count() *int64 {
	return &counter.count
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var wc WriterCounter
	wc.writer = w
	return &wc, wc.Count()
}

// atomic.LoadUint64(&counter.count)
// atomic.AddUint64(&counter.count, uint64(n))