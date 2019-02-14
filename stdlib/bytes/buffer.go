package bytes

// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
	buf       []byte
	off       int
	bootstrap [64]byte
	lastRead  readOp
}
