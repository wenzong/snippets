#include <stdio.h>
#include <assert.h>
#include <uv.h>

uv_fs_t open_req;
uv_fs_t read_req;
uv_fs_t write_req;
static char buffer[1024];
static uv_buf_t iov;

// uv_fs_t - Request Type
//
// + Handles represent long-lived objects.
// + Async operations on such handles are identified using requests
//
// typedef void (*uv_fs_cb)(uv_fs_t* req);
void on_open(uv_fs_t *req);
void on_read(uv_fs_t *req);
void on_write(uv_fs_t *req);

void on_open(uv_fs_t *req) {
    assert(req == &open_req);
    if (req->result >= 0) {
        iov = uv_buf_init(buffer, sizeof(buffer));
        // int uv_fs_read(uv_loop_t* loop, uv_fs_t* req, uv_file file, const uv_buf_t bufs[], unsigned int nbufs, int64_t offset, uv_fs_cb cb)
        // ssize_t preadv(                                     int fd, const struct iovec *iov, int iovcnt, off_t offset);
        uv_fs_read(uv_default_loop(), &read_req, req->result, &iov, 1, -1, on_read);
    } else {
        fprintf(stderr, "Open error: %s\n", uv_strerror(req->result));
    }
}

void on_read(uv_fs_t *req) {
    assert(req == &read_req);
    if (req->result < 0) {
        fprintf(stderr, "Read error: %s\n", uv_strerror(req->result));
    }
    else if (req->result == 0) {
        uv_fs_t close_req;
        // synchronous
        uv_fs_close(uv_default_loop(), &close_req, open_req.result, NULL);
    }
    else if (req->result > 0) {
        iov.len = req->result;
        // int uv_fs_write(uv_loop_t* loop, uv_fs_t* req, uv_file file, const uv_buf_t bufs[], unsigned int nbufs, int64_t offset, uv_fs_cb cb)
        // ssize_t pwritev(                                     int fd, const struct iovec *iov, int iovcnt, off_t offset);
        uv_fs_write(uv_default_loop(), &write_req, 1/* stdout */, &iov, 1, -1, on_write);
    }
}

void on_write(uv_fs_t *req) {
    if (req->result < 0) {
        fprintf(stderr, "Write error: %s\n", uv_strerror((int)req->result));
    }
    else {
        uv_fs_read(uv_default_loop(), &read_req, open_req.result, &iov, 1, -1, on_read);
    }
}

int main(int argc, char **argv) {
    if (argc != 2) {
        fprintf(stderr, "usage: ./cat /PATH/TO/FILE\n");
        return 1;
    }
    // int uv_fs_open(uv_loop_t* loop, uv_fs_t* req, const char* path, int flags, int mode, uv_fs_cb cb)
    // int open(                                     const char *path, int oflag, ...);
    uv_fs_open(uv_default_loop(), &open_req, argv[1], O_RDONLY, 0, on_open);

    uv_run(uv_default_loop(), UV_RUN_DEFAULT);

    uv_fs_req_cleanup(&open_req);
    uv_fs_req_cleanup(&read_req);
    uv_fs_req_cleanup(&write_req);

    return 0;
}
