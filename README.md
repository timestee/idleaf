# idleaf  [![Build Status](https://travis-ci.org/timestee/idleaf.svg?branch=master)](https://travis-ci.org/timestee/idleaf) 
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
