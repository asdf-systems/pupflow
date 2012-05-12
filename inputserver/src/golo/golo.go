package golo

// #cgo pkg-config: liblo
// #include <lo/lo.h>
// #include <lo/lo.h>
// #include "golo.h"
import "C"

import (
	"unsafe"
	"errors"
)

type Message struct {
	Path string
	Params []interface{}
}

func Deserialize(data []byte) (*Message, error) {
	cdata := unsafe.Pointer(&data[0])
	cdatalen := C.size_t(len(data))

	msg := C.lo_message_deserialise(cdata, cdatalen, (*C.int)(unsafe.Pointer(nil)))
	if msg == nil {
		// TODO: Obtain and parse actual error?
		return nil, errors.New("Deserialisation failed")
	}

	result := &Message{
		Path: C.GoString(C.lo_get_path(cdata, C.ssize_t(cdatalen))),
		Params: make([]interface{}, 0, 10),
	}
	argc := int(C.lo_message_get_argc(msg))
	argtypes := C.GoString(C.lo_message_get_types(msg))
	argv := C.lo_message_get_argv(msg)
	for i := 0; i < argc; i++ {
		var val interface{}
		switch argtypes[i] {
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
		result.Params = append(result.Params, val)
	}
	return result, nil
}
