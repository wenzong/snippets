## TL;DR

```shell
$ make serve
```

## Problems

### protobuf int64/google.protobuf.Timestamp

+ sql: Scan error on column index 6, name "created_at": converting driver.Value type time.Time ("2020-07-03 15:41:04 +0000 UTC") to a int64: invalid syntax
+ sql: Scan error on column index 6, name "created_at": unsupported Scan, storing driver.Value type time.Time into type *timestamppb.Timestamp
