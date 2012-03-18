package golo

// #cgo pkg-config: liblo
// #include <lo/lo.h>
// #include <lo/lo.h>
// #include "golo.h"
import "C"

import (
	// "fmt"
	"unsafe"
)

var (
	chanmap = make(map[int]chan *Message)
)

type Message struct {
	Path string
	Params []interface{}
}

func StartServer(port int, c chan *Message) {
	C.start_server(C.int(port))
	chanmap[port] = c
}

//export callback
func callback(path, ctypes *C.char, argv **C.lo_arg, argc int, user_data unsafe.Pointer) {
	panic("IN HERE")
	c, ok := chanmap[int(uintptr(user_data))]
	if !ok {
		panic("Unknown callback")
	}

	msg := &Message {
		Path: C.GoString(path),
		Params: make([]interface{}, 0, 1),
	}
	types := C.GoString(ctypes)
	for i := 0; i < argc; i++ {

		val := interface{}(nil)
		switch types[i] {
		case 'i':
			val = int32(C.msg_extract_int32(argv, C.int(i)))
		case 'h':
			val = int64(C.msg_extract_int64(argv, C.int(i)))
		case 'f':
			val = float32(C.msg_extract_float32(argv, C.int(i)))
		case 'd':
			val = float64(C.msg_extract_float64(argv, C.int(i)))
		default:
			val = "Unknown Type"
		}
		msg.Params = append(msg.Params, val)
	}
	c <- msg
}
