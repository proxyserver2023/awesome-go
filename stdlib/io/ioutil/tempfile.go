// package ioutil implements some I/O utility functions.
package ioutil

import "io"

// readAll reads from r until an error or EOF and returns the data
// it read
// from the internal buffer
// allocated with a specific capacity
func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	return
}
