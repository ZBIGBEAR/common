package translate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslate(t *testing.T) {
	source := "你好"
	trans := New()
	result, err := trans.Translate(source, ZH, EN)
	assert.Nil(t, err)
	assert.Equal(t, "Hello", result)
}