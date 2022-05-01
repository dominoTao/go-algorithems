package main


import "io"

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int64, error) {
	*b += ByteCounter(len(p))
	return int64(len(p)), nil
}

type countingWriter struct {
	count  *int64
	writer io.Writer
}

func (receiver *countingWriter) Write(a []byte) (int, error) {
	n, err := receiver.writer.Write(a)
	if err == nil {
		*receiver.count += int64(n)
	}
	return n, err
}

// CountingWriter 返回Writer接口和写入数据的字节数指针
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var a countingWriter
	var n int64
	a.writer = w
	a.count = &n
	return a.writer, a.count
}

