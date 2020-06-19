package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	Data unsafe.Pointer
	Len int
}

type SliceHeader struct {
	Data unsafe.Pointer
	Len int
	Cap int
}

func strLen(s string) int {
	header := (*StringHeader)(unsafe.Pointer(&s))
	return header.Len
}

func strData(s string) uintptr {
	header := (*StringHeader)(unsafe.Pointer(&s))
	return uintptr(header.Data)
}

func strToBytes(s string) []byte {
	header := (*StringHeader)(unsafe.Pointer(&s))
	bytesHeader := &SliceHeader{
		Data: header.Data,
		Len: header.Len,
		Cap: header.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bytesHeader))
}

func bytesToStr(b []byte) string {
	header := (*SliceHeader)(unsafe.Pointer(&b))
	strHeader := &StringHeader{
		Data: header.Data,
		Len: header.Len,
	}
	return *(*string)(unsafe.Pointer(strHeader))
}

func setStrLen(s string, length int) {
	header := (*StringHeader)(unsafe.Pointer(&s))
	header.Len = length
}

func withNewStrLen(s string, length int) string {
	header := (*StringHeader)(unsafe.Pointer(&s))
	header.Len = length
	return *(*string)(unsafe.Pointer(header))
}

func substr(s string, start, length int) string {
	header := (*StringHeader)(unsafe.Pointer(&s))
	i := uintptr(start)
	header.Data = unsafe.Pointer(uintptr(header.Data) + i*unsafe.Sizeof(s[0]))
	header.Len = length
	return *(*string)(unsafe.Pointer(header))
}

func main() {
	fmt.Printf("%d\n", strLen("asdfasdfasdf"))
	fmt.Printf("%v\n", strToBytes("asdfasdfasdf"))
	fmt.Printf("%v\n", bytesToStr([]byte{97, 115, 100, 102, 97, 115, 100, 102, 97, 115, 100, 102,}))

	s0 := "asdfasdf"
	b0 := strToBytes(s0)
	s1 := bytesToStr(b0)

	// b0[0] = 98
	//> panic: runtime error: invalid memory address or nil pointer dereference
	//> [signal SIGSEGV: segmentation violation code=0xffffffff addr=0x0 pc=0xd56a2]

	fmt.Printf("strLen(s0): %d\n", strLen(s0))
	fmt.Printf("strData(s0): %v\n", strData(s0))
	fmt.Printf("s0: %v\n", s0)
	fmt.Printf("b0: %v\n", b0)
	fmt.Printf("s1: %v\n", s1)

	b1 := []byte{97, 115, 100, 102, 97, 115, 100, 102}
	s2 := bytesToStr(b1)
	b2 := strToBytes(s2)

	b1[0] = 98
	b1[5] = 0

	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("b2: %v\n", b2)

	setStrLen(s0, 128)
	fmt.Printf("extended s0 at %v: %v\n", strData(s0), s0)
	s3 := withNewStrLen(s0, 128)
	fmt.Printf("really extended s0 at %v: %v\n", strData(s3), s3)
	fmt.Printf("substr s0: %v\n", substr(s0, 2, 5))
}
