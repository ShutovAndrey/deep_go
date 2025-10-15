package main

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type Numeric interface {
	uint32 | uint16 | uint64
}

func ToLittleEndian[T Numeric](number T) T {
	size := int(unsafe.Sizeof(number))
	var result T
	for i := 0; i < size; i++ {
		result |= (number >> (8 * uint(i))) & 0xFF << (8 * uint(size-1-i))
	}
	return result
}

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number any
		result any
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
		"test case #6": {
			number: 0x00FF,
			result: 0xFF00,
		},
		"test case #7": {
			number: 0x0102030405060708,
			result: 0x0807060504030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			switch n := test.number.(type) {
			case uint16:
				assert.Equal(t, test.result.(uint16), ToLittleEndian(n))
			case uint32:
				assert.Equal(t, test.result.(uint32), ToLittleEndian(n))
			case uint64:
				assert.Equal(t, test.result.(uint64), ToLittleEndian(n))
			}
		})
	}
}
