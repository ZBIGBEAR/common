package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input interface{}
	want  bool
}

func TestIsEmpty(t *testing.T) {
	// 结构体指针对象为空
	var tc1 *TestCase
	result := IsEmpty(tc1)
	assert.True(t, result)

	// 结构体对象不为空
	tc2 := TestCase{}
	result = IsEmpty(tc2)
	assert.False(t, result)

	// 结构体指针对象不为空
	result = IsEmpty(&tc2)
	assert.False(t, result)

	// 结构体指针为空
	tc3 := &tc2
	tc3 = nil
	result = IsEmpty(tc3)
	assert.True(t, result)

	// // 数组未初始化
	var tcs []TestCase
	result = IsEmpty(tcs)
	assert.True(t, result)

	// 数组已初始化，但是没有元素
	tcs = []TestCase{}
	result = IsEmpty(tcs)
	assert.True(t, result)

	// 未初始化的map
	var m map[string]string
	result = IsEmpty(m)
	assert.True(t, result)

	// 已初始化的空map
	m = make(map[string]string)
	result = IsEmpty(m)
	assert.True(t, result)

	// 不为空的map
	m["a"] = "b"
	result = IsEmpty(m)
	assert.False(t, result)

	var b bool
	var bp *bool
	i := int(0)
	var ip *int
	e1 := uint(0)
	var p1 *uint
	e2 := float32(0)
	var p2 *float32
	e3, v3 := "", "test"
	var p3 *string

	f1 := float64(0.00)
	testCases := []TestCase{
		{input: b, want: false},
		{input: &b, want: false},
		{input: bp, want: true},
		{input: i, want: true},
		{input: &i, want: true},
		{input: int(1), want: false},
		{input: ip, want: true},
		{input: e1, want: true},
		{input: &e1, want: true},
		{input: uint(1), want: false},
		{input: p1, want: true},
		{input: e2, want: true},
		{input: &e2, want: true},
		{input: float32(1), want: false},
		{input: p2, want: true},
		{input: e3, want: true},
		{input: &e3, want: true},
		{input: v3, want: false},
		{input: p3, want: true},
		{input: f1, want: true},
		{input: &f1, want: true},
	}

	for i, testCase := range testCases {
		result = IsEmpty(testCases[i].input)
		assert.Equal(t, testCase.want, result)
	}
}

func BenchmarkTestIsEmpty(t *testing.B) {
	// 结构体指针对象为空
	var tc1 *TestCase

	// 结构体对象不为空
	tc2 := TestCase{}

	// 结构体指针为空
	tc3 := &tc2
	tc3 = nil

	for i := 0; i < t.N; i++ {
		result := IsEmpty(tc1)
		assert.True(t, result)

		// 结构体指针对象不为空
		result = IsEmpty(&tc2)
		assert.False(t, result)

		result = IsEmpty(tc3)
		assert.True(t, result)
	}
}

func BenchmarkTestIsEmpty1(t *testing.B) {
	// 结构体指针对象为空
	var tc1 *TestCase

	// 结构体对象不为空
	tc2 := TestCase{}

	// 结构体指针为空
	tc3 := &tc2
	tc3 = nil

	for i := 0; i < t.N; i++ {
		result := tc1 == nil
		assert.True(t, result)

		// 结构体指针对象不为空
		v := &tc2
		result = v == nil
		assert.False(t, result)

		result = tc3 == nil
		assert.True(t, result)
	}
}
