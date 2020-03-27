# lru
lru cache

![Build and Test](https://github.com/two/lru/workflows/Build%20and%20Test/badge.svg)
[![Code Coverage](https://codecov.io/gh/two/lru/branch/master/graph/badge.svg)](https://codecov.io/gh/two/lru)
[![GoDoc](https://godoc.org/github.com/two/lru?status.svg)](https://godoc.org/github.com/two/lru)

## 使用方法

```go
lru := NewLRU()


## set max capacity 
lru.SetMaxCap(10)

## set key value
lru.Set("key1", "value1")
lru.Set("key2", "value2")

## get key value
val := lru.Get("key1").(string)
println(val)

## remove one key
lru.Remove("key1")

## clear lru cache
lru.Clear()
```

