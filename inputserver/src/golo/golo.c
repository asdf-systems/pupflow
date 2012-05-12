#include <stdio.h>
#include <stdint.h>
#include <lo/lo.h>
#include "golo.h"
#include "_cgo_export.h"

int32_t msg_extract_int32(lo_arg **argv, int pos) {
	return argv[pos]->i32;
}

int64_t msg_extract_int64(lo_arg **argv, int pos) {
	return argv[pos]->i64;
}

float msg_extract_float32(lo_arg **argv, int pos) {
	return argv[pos]->f32;
}

double msg_extract_float64(lo_arg **argv, int pos) {
	return argv[pos]->f64;
}

