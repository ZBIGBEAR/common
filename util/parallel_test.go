package util

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParallel(t *testing.T) {
	parallel := NewDefaultParallel()

	parallel.Add(func() error {
		return errors.New("tets err 1")
	})
	parallel.Add(func() error {
		return errors.New("tets err 2")
	})

	parallel.Add(func() error {
		return nil
	})

	parallel.Start()

	result := parallel.Result()
	fmt.Println(result)
	assert.Equal(t, 2, len(result))
}

func TestParallel1(t *testing.T) {
	parallel := NewParallel(5)
	for i := 0; i < 10; i++ {
		parallel.Add(func() error {
			time.Sleep(time.Second)
			return nil
		})
	}

	parallel.Start()
	result := parallel.Result()
	assert.Equal(t, 0, len(result))
}
