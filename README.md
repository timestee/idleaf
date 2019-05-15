# idleaf [![Build Status](https://travis-ci.org/timestee/idleaf.svg?branch=master)](https://travis-ci.org/timestee/idleaf) [![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/timestee/idleaf)  [![GoDoc](https://godoc.org/github.com/timestee/idleaf?status.svg)](https://godoc.org/github.com/timestee/idleaf)

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
