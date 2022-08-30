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
	v1 := false
	v2 := true
	var v3 *bool
	v4 := int(0)
	var v5 *int
	v6 := int64(0)
	var v7 *int64
	v8 := ""
	var v9 *string
	v10 := "test"
	v11 := float64(0.00)
	v12 := float64(0.01)

	testCases := []TestCase{
		{input: v1, want: false},
		{input: &v1, want: false},
		{input: v2, want: false},
		{input: &v2, want: false},
		{input: v3, want: true},
		{input: v4, want: true},
		{input: &v4, want: true},
		{input: v5, want: true},
		{input: int(1), want: false},
		{input: v6, want: true},
		{input: &v6, want: true},
		{input: v7, want: true},
		{input: int64(1), want: false},
		{input: v8, want: true},
		{input: &v8, want: true},
		{input: v9, want: true},
		{input: v10, want: false},
		{input: &v10, want: false},
		{input: v11, want: true},
		{input: &v11, want: true},
		{input: v12, want: false},
		{input: &v12, want: false},
	}

	for i, testCase := range testCases {
		result := IsEmpty(testCases[i].input)
		assert.Equal(t, testCase.want, result)
	}
}

func BenchmarkTestIsEmptyV1(t *testing.B) {
	v1 := false
	v2 := true
	var v3 *bool
	v4 := int(0)
	var v5 *int
	v6 := int64(0)
	var v7 *int64
	v8 := ""
	var v9 *string
	v10 := "test"
	v11 := float64(0.00)
	v12 := float64(0.01)

	testCases := []TestCase{
		{input: v1, want: false},
		{input: &v1, want: false},
		{input: v2, want: false},
		{input: &v2, want: false},
		{input: v3, want: true},
		{input: v4, want: true},
		{input: &v4, want: true},
		{input: v5, want: true},
		{input: int(1), want: false},
		{input: v6, want: true},
		{input: &v6, want: true},
		{input: v7, want: true},
		{input: int64(1), want: false},
		{input: v8, want: true},
		{input: &v8, want: true},
		{input: v9, want: true},
		{input: v10, want: false},
		{input: &v10, want: false},
		{input: v11, want: true},
		{input: &v11, want: true},
		{input: v12, want: false},
		{input: &v12, want: false},
	}

	for i, testCase := range testCases {
		result := IsEmpty(testCases[i].input)
		assert.Equal(t, testCase.want, result)
	}
}

/*
压测结果如下
# go test -bench=.  -benchmem
result:
BenchmarkTestIsEmpty-12         34300347                34.67 ns/op            0 B/op          0 allocs/op
BenchmarkTestIsEmpty1-12        204903328                5.938 ns/op           0 B/op          0 allocs/op

IsEmpty里面用了反射，跟不用反射相比，性能相差将近6倍
*/
