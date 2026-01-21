package main

/*
#cgo CFLAGS: -mavx2
#include <immintrin.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	buf := [32]uint8{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F, 0x20,
	}
	// reflect.TypeOf(&buf) is *[32]uint8
	ptr := (*C.__m256i)(unsafe.Pointer(&buf)) // go from *[32]uint8 to *C.__m256i that are of same size, but *C.__m256i has alignment of 32
	x := *ptr                                 // can spuriously panic due to misalignment
	fmt.Printf("Loaded value: %v\n", x)
}
