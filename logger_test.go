package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {

	condition := func(level int) bool {
		if level >= 3 {
			return true
		}
		return false
	}

	logger := NewBuilder().WriteToConsole().WithTimestamp().If(condition).Build()

	logger.Log(1, "1")
	logger.Log(2, "2")
	logger.Log(3, "3")
	logger.Log(4, "4")

	assert.Equal(t, 123, 123, "they should be equal")
}
