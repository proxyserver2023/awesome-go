package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError(t *testing.T) {
	t.Run("test error method", func(t *testing.T) {
		e := APIError{
			Message: "abc",
		}
		assert.Equal(t, "abc", e.Error())
	})

	t.Run("test status code", func(t *testing.T) {
		e := APIError{
			Status: 404,
		}
		assert.Equal(t, e.Status, e.StatusCode())
	})
}
