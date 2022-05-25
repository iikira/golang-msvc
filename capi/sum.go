package main

//#include <string.h>
import "C"
import (
	"crypto/md5"
	"encoding/hex"
	"unsafe"
)

//export add
func add(a, b C.int) C.int {
	return C.int(int(a) + int(b))
}

//export md5sum
func md5sum(str *C.char, res *C.char) {
	m := md5.New()
	m.Write(C.GoBytes(unsafe.Pointer(str), C.int(C.strlen(str))))
	h := make([]byte, 33)
	hex.Encode(h, m.Sum(nil))
	C.memcpy(unsafe.Pointer(res), unsafe.Pointer(&h[0]), 33)
}
