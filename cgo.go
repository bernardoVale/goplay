package main

/*
//#include <stdlib.h>
*/

import "C"
import (
	"fmt"
	"unsafe"
)

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	fmt.Printf("%+v", cs)
	//C.fputs(cs, (*C.FILE)(C.stdout))

}
func main() {
	Print("Hello C")
	fmt.Print("%+v")
}
