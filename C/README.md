```c
// https://stackoverflow.com/a/24953446
void (*signal(int sig, void (*func)(int)))(int)

// or
typedef void (*sighandler_t)(int)
sighandler_t signal(int sig, sighandler_t func);
```

```go
type func signal(int, func(int)) func(int)

// or

type sighandler_t func(int)
type signal func(int, sighandler_t) sighandler_t
```
