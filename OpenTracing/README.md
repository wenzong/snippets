```
$ docker run --rm -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest
$ go run main.go
```

```
$ curl http://127.0.0.1:8080/hello
$ open http://127.0.0.1:16681
```
