package bytes_test

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
