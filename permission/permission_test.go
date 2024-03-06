package permission

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerm(t *testing.T) {
	p := NewPerm()
	p.AddPerm(CreatePerm)
	p.AddPerm(ReadPerm)
	assert.True(t, p.HasPerm(ReadPerm))
	assert.False(t, p.HasPerm(WritePerm))

	p.RemovePerm(ReadPerm)
	assert.False(t, p.HasPerm(ReadPerm))
}
