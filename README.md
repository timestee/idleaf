Note: This project is no longer actively maintained. Replaced by siid for much better performance(funplus internal use only now).

siid Benchmark Results

With Go 1.13.4 darwin/amd64 on a 2.2 GHz Intel Core i7 16GB. Like all benchmarks, take these with a grain of salt.
```shell

goos: darwin
goarch: amd64
pkg: bitbucket.org/funplus/siid/bentch
BenchmarkSIID_Mongo-12      46535821            25.4 ns/op         0 B/op          0 allocs/op
BenchmarkSIID_MySQL-12      40794764            27.4 ns/op         0 B/op          0 allocs/op
BenchmarkRand-12            92630166            12.9 ns/op         0 B/op          0 allocs/op
BenchmarkTimestamp-12       17214086            69.3 ns/op         0 B/op          0 allocs/op
BenchmarkUUID_V1-12         11488557           107 ns/op           0 B/op          0 allocs/op
BenchmarkUUID_V2-12         11898156           101 ns/op           0 B/op          0 allocs/op
BenchmarkUUID_V3-12          5220328           229 ns/op         144 B/op          4 allocs/op
BenchmarkUUID_V4-12         15463155            76.4 ns/op        16 B/op          1 allocs/op
BenchmarkUUID_V5-12          4481305           271 ns/op         176 B/op          4 allocs/op
BenchmarkSnowflake-12        4928191           244 ns/op           0 B/op          0 allocs/op
PASS
ok      bitbucket.org/funplus/siid/bentch   14.801s
```


# idleaf
[![Build Status](https://travis-ci.org/timestee/idleaf.svg?branch=master)](https://travis-ci.org/timestee/idleaf) 
[![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/timestee/idleaf)  
[![GoDoc](https://godoc.org/github.com/timestee/idleaf?status.svg)](https://godoc.org/github.com/timestee/idleaf)
[![Go Report Card](https://goreportcard.com/badge/github.com/timestee/idleaf)](https://goreportcard.com/report/github.com/timestee/idleaf)

Generate `integer` id through MySQL transaction way, dose not increase the load of MySQL.

Usage:
```shell
curl http://127.0.0.1:16000/v1/gen/user_id
```

Response:
```json
{"code":0,"msg":"","id":1000001}
```

## Architecture overview
```
+---------+    +---------+  ...  +---------+    +---------+
| service |    | service |  ...  | service |    | service |
+---------+    +---------+  ...  +---------+    +---------+
         \       /     \         /     |       /
          \     /       \       /      |      /
	  +--------+    +--------+    +--------+
	  | idleaf |    | idleaf |    | idleaf |
	  +--------+    +--------+    +--------+
                  \       |         /
                   \      |        /
                    \     |       /
                    +----+------+
                         |
                         DB
```
