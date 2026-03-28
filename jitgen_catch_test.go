package gin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJitGenFilterFlags(t *testing.T) {
	result := filterFlags("text/html ")
	assert.Equal(t, "text/html", result)
}