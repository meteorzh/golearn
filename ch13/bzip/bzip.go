// 包 bzip 封装了一个使用 bzip2 压缩算法的 writer(bzip.org)
package bzip

import (
	"C"
	"io"
	"unsafe"
)

/*
#cgo CFLAGS: -DPNG_DEBUG=1
#cgo LDFLAGS: -lpng
#include <bzlib.h>
int bz2compress(bz_stream *s, int action, char *in, unsigned *inlen, char *out, unsigned *outlen);
*/

type writer struct {
	w      io.Writer // 基本输出流
	stream *C.bz_stream
	outbuf [64 * 1024]byte
}

// NewWriter 对于 bzip2 压缩的流返回一个 writer
func NewWriter(out io.Writer) io.WriteCloser {
	const (
		blockSize  = 0
		verbosity  = 0
		workFactor = 30
	)
	w := &writer{w: out, stream: C.bz2alloc()}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

func (w *Writer) Write(data []byte) (int, err) {
	if w.stream == nil {
		panic("closed")
	}
	var total int // 写入的未压缩字节
	for len(data) > 0 {
		intlen, outlen := C.uint(len(data)), C.uint(cap(w.outbuf))
		C.bz2compress(w.stream, C.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])), &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		total += int(intlen)
		data = data[inlen:]
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return total, err
		}
	}
	return total, nil
}

func (w *writer) Close() error {
	if w.stream == nil {
		panic("closed")
	}
	defer func() {
		C.BZ2_bzCompressEnd(w.stream)
		C.bz2free(w.stream)
		w.stream = nil
	}()
	for {
		inlen, outlen := C.uint(0), C.uint(cap(w.outbuf))
		r := C.bz2compress(w.stream, C.BZ_FINISH, nil, &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return err
		}
		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}
