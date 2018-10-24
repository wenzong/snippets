/*
 * man 3 signal(BSD Library Functions)
 *
 * #include <signal.h>
 * void (*signal(int sig, void (*func)(int)))(int);
 *
 * or
 *
 * typedef void (*sig_t) (int);
 *
 * sig_t
 * signal(int sig, sig_t func);
 *
 */

// otherwise assert won't work
#undef NDEBUG

#include <assert.h>
#include <signal.h>
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

#include "dbg.h"

static void sig_handler(int);

int
main() {
    sig_t prev_sig_handler;

    if ((prev_sig_handler = signal(SIGINT, sig_handler)) == SIG_ERR) {
        log_err("signal error");
    }

    assert(prev_sig_handler == NULL);

    if ((prev_sig_handler = signal(SIGINT, sig_handler)) == SIG_ERR) {
        log_err("signal error");
    }

    assert(prev_sig_handler == sig_handler);

    if (signal(SIGSTOP, sig_handler) == SIG_ERR) {
        log_err("signal error");
    }

    while (1) {
        log_info("I am pid<%i> process", getpid());
        sleep(1);
    }
    return 0;
}

void
sig_handler(int signo) {
    log_info("Hello Signal: %i", signo);
}
