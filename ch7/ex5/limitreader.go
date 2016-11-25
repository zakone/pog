package limitreader

import (
    "io"
)

type LimitedReader struct {
    limit  int64
    reader io.Reader
}

func (h *LimitedReader) Read(b []byte) (n int, err error) {
    if int64(len(b)) > h.limit {
        b = b[:h.limit]
    }
    n, err = h.reader.Read(b[:h.limit])
    return
}

func LimitReader(r io.Reader, n int64) io.Reader {
    return &LimitedReader{n, r}

}
