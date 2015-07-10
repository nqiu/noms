package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimitiveEquals(t *testing.T) {
	assert := assert.New(t)

	values := []func() Value{
		func() Value { return nil },
		func() Value { return Bool(false) },
		func() Value { return Bool(true) },
		func() Value { return Int16(0) },
		func() Value { return Int16(1) },
		func() Value { return Int16(-1) },
		func() Value { return Int32(0) },
		func() Value { return Int32(1) },
		func() Value { return Int32(-1) },
		func() Value { return Int64(0) },
		func() Value { return Int64(1) },
		func() Value { return Int64(-1) },
		func() Value { return UInt16(0) },
		func() Value { return UInt16(1) },
		func() Value { return UInt32(0) },
		func() Value { return UInt32(1) },
		func() Value { return UInt64(0) },
		func() Value { return UInt64(1) },
		func() Value { return Float32(0) },
		func() Value { return Float32(-1) },
		func() Value { return Float32(1) },
		func() Value { return Float64(0) },
		func() Value { return Float64(-1) },
		func() Value { return Float64(1) },
		func() Value { return NewString("") },
		func() Value { return NewString("hi") },
		func() Value { return NewString("bye") },
		func() Value { return NewBlob([]byte{}) },
		func() Value { return NewBlob([]byte("hi")) },
		func() Value { return NewBlob([]byte("bye")) },
		func() Value { return NewList() },
		func() Value { return NewList(NewString("foo")) },
		func() Value { return NewList(NewString("bar")) },
		func() Value { return NewMap() },
		func() Value { return NewMap(NewString("a"), NewString("a")) },
		func() Value { return NewSet() },
		func() Value { return NewSet(NewString("hi")) },
	}

	for i, f1 := range values {
		for j, f2 := range values {
			if f1() == nil {
				continue
			}
			if i == j {
				assert.True(f1().Equals(f2()))
			} else {
				assert.False(f1().Equals(f2()))
			}
		}
	}
}