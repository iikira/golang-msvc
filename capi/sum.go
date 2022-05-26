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
