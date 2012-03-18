#include <stdio.h>
#include <stdint.h>
#include <lo/lo.h>
#include "golo.h"
#include "_cgo_export.h"

void error(int num, const char *msg, const char *path) {
    printf("liblo server error %d in path %s: %s\n", num, path, msg);
    fflush(stdout);
}

int generic_handler(const char *path, const char *types, lo_arg **argv, int argc, void *data, void *user_data) {
	fprintf(stderr, "C IN HERE\n");
	callback(path, types, argv, argc, user_data);
	return 1;
}

void start_server(int port) {
	char addr[7];
	sprintf(addr, "%d", port);
	lo_server_thread s = lo_server_thread_new(addr, error);
	lo_server_thread_add_method(s, NULL, NULL, generic_handler, (void*)port);
	lo_server_thread_start(s);
}

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

