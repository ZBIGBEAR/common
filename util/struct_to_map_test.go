package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Name    string `json:" name  ,omitempty" structs:"name"`
	Age     int    `json:"age,omitempty" structs:"age,omitempty"`
	IsMarry bool   `json:"is_marry,omitempty" structs:"is_marry"`
}

func TestStructToMap(t *testing.T) {
	testS := &testStruct{
		Name: "abc",
		Age:  0,
	}

	m1, err := StructToMapV1(testS)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(m1))

	m2, err := StructToMapV2(testS)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(m2))

	m3 := StructToMapV3(testS)
	assert.Equal(t, 2, len(m3))
}

func BenchmarkStructToMapV1(t *testing.B) {
	testS := &testStruct{
		Name: "abc",
		Age:  0,
	}

	for i := 0; i < t.N; i++ {
		StructToMapV1(testS)
	}
}

func BenchmarkStructToMapV2(t *testing.B) {
	testS := &testStruct{
		Name: "abc",
		Age:  0,
	}

	for i := 0; i < t.N; i++ {
		StructToMapV2(testS)
	}
}

func BenchmarkStructToMapV3(t *testing.B) {
	testS := &testStruct{
		Name: "abc",
		Age:  0,
	}

	for i := 0; i < t.N; i++ {
		StructToMapV3(testS)
	}
}

/*
BenchmarkStructToMapV1-12        1000000              1015 ns/op
BenchmarkStructToMapV2-12        1206841               989.7 ns/op
BenchmarkStructToMapV3-12         588036              1792 ns/op

压测显示V2运行最快
*/
