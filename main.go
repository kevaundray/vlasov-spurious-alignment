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
	buf := [32]uint8{}
	// reflect.TypeOf(&buf) is *[32]uint8
	ptr := (*C.__m256i)(unsafe.Pointer(&buf)) // go from *[32]uint8 to *C.__m256i that are of same size, but *C.__m256i has alignment of 32
	x := *ptr                                 // can spuriously panic due to misalignment
	fmt.Printf("Loaded value: %v\n", x)
}
