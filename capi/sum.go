package main

//#include <string.h>
import "C"
import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"unsafe"
)

//export add
func add(a, b C.int) C.int {
	return C.int(int(a) + int(b))
}

//export go_complex64
func go_complex64(a, b float32) complex64 {
	// gcc 和 msvc 初始化complex的方式不一样, 故统一使用go来初始化complex
	return complex(a, b)
}

//export go_complex128
func go_complex128(a, b float64) complex128 {
	return complex(a, b)
}

//export complex64_add
func complex64_add(a, b complex64) complex64 {
	return a + b
}

//export complex128_add
func complex128_add(a, b complex128) complex128 {
	return a + b
}

// res must allocate size more than 33
//export md5sum
func md5sum(str *C.char, res *C.char) {
	m := md5.New()
	l := int(C.strlen(str))
	m.Write(*(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(str)),
		Len:  l,
		Cap:  l,
	})))
	h := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(res)),
		Len:  33,
		Cap:  33,
	}))
	hex.Encode(h, m.Sum(nil))
	h[32] = 0
}
