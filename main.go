package main

/*
#cgo CFLAGS: -mavx2
#include <immintrin.h>
*/
import "C"
import (
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

//go:noinline
func testAlignment(depth int) {
	// Each recursion level shifts the stack
	if depth > 0 {
		var pad [1]byte // 1-byte pad to shift stack alignment
		_ = pad
		testAlignment(depth - 1)
		return
	}

	buf := [32]uint8{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0x20,
	}

	addr := uintptr(unsafe.Pointer(&buf))
	fmt.Printf("depth=%d, addr=%#x, mod32=%d\n", depth, addr, addr%32)

	ptr := (*C.__m256i)(unsafe.Pointer(&buf))
	x := *ptr // can panic if misaligned
	fmt.Printf("Loaded value: %v\n", x)
}

func main() {
	depth := 0
	if len(os.Args) > 1 {
		depth, _ = strconv.Atoi(os.Args[1])
	}
	testAlignment(depth)
}
