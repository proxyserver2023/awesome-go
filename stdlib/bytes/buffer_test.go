package bytes_test

import "testing"

const N = 1000

var testString string
var testBytes []byte

type negativeReader struct{}

func (r *negativeReader) Read([]byte) (int, error) {
	return -1, nil
}

func init() {
	testBytes = make([]byte, N)
	for i := 0; i < N; i++ {
		testBytes[i] = 'a' + byte(i%26)
	}
	testString = string(testBytes)
}

// Verify that contents of buf match the string s
func check(t *testing.T, testname string, buf *Buffer, s string) {
	bytes := buf.Bytes()
	str := buf.String()

	if buf.Len() != len(bytes) {
		t.Errorf("%s: buf.Len() == %d, len(buf.Bytes()) == %d", testname, buf.Len(), len(bytes))
	}

	if buf.Len() != len(str) {
		t.Errorf("%s: buf.Len() == %d, len(buf.Strings()) == %d", testname, buf.Len(), len(str))
	}

	if buf.Len() != len(s) {
		t.Errorf("%s: buf.Len() == %d, len(s) == %d", testname, buf.Len(), len(s))
	}

	if string(bytes) != s {
		t.Errorf("%s: string(buf.Bytes()) == %s, s == %s", testname, string(bytes), s)
	}

}
