package main

import (
	"fmt"
	"unsafe"
)

type ifaceInternal struct {
	p1 uintptr
	p2 uintptr
}

var table interface{}

func main() {
	var iface interface{} = byte(5)
	internal := *(*ifaceInternal)(unsafe.Pointer(&iface))
	fmt.Printf("%x %x\n", internal.p1, internal.p2)
	printBytes()
}

func printBytes() {
	for i := 0; i < 10; i++ {
		b := byte(i)
		table = &b
		bpp := table.(*byte)
		bp := uintptr(unsafe.Pointer(bpp))
		fmt.Printf("%d = %x\n", i, bp)
	}
}