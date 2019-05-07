# idleaf [![Build Status](https://travis-ci.org/timestee/idleaf.svg?branch=master)](https://travis-ci.org/timestee/idleaf) [![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/timestee/idleaf)  [![GoDoc](https://godoc.org/github.com/timestee/idleaf?status.svg)](https://godoc.org/github.com/timestee/idleaf)

Id generator for golang. There are no two identical leaves in the world.

Generate batch ids through MySQL transcation way, dose not increase the load of MySQL.

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
