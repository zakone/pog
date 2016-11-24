package limitReader

import (
    "fmt"
    "io"
    "os"
)

type LimitedReader struct {
    limit  int64
    reader io.Reader
}

func (h *LimitReader) Read(b []byte) (n int, err error) {
    if len(b) > h.limit {
        b = b[:h.limit]
    }
    n, err := h.reader.Read(b[:h.limit])
    return
}

func LimitReader(r io.Reader, n int64) io.Reader {
    return &LimitedReader{r, n}

}
