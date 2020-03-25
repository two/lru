package lru

import (
	"sync"
)

// LRU 定义 LRU 接口
type LRU interface {
	Get(key string) interface{}      // 通过 key 获取缓存
	Set(key string, val interface{}) // 设置key 及其对应的值
	//SetWithExpire(key string, t time.Duration, val interface{}) // 设置缓存，同时设置存活时间
	//SetWithExpireAt(key string, t time.Time, val interface{})   // 设置缓存，同时设置过期时间
	Remove(key string) // 删除某个 key 对应的缓存
	Clear()            // 清空所有缓存
	MaxCap() int
	SetMaxCap(int)
}

var defaultLRUCap = 2 ^ 10

type lruVal struct {
	val  interface{}
	node *Node
}

type defaultLRU struct {
	m      map[string]lruVal
	list   List
	maxCap int
	sync.Mutex
}

// NewLRU create an defalt lru object
func NewLRU() LRU {
	return &defaultLRU{
		m:      make(map[string]lruVal),
		list:   List{},
		maxCap: defaultLRUCap,
	}
}

func (lru *defaultLRU) MaxCap() int {
	return lru.maxCap
}

func (lru *defaultLRU) SetMaxCap(maxCap int) {
	lru.maxCap = maxCap
}

func (lru *defaultLRU) Get(key string) interface{} {
	lru.Lock()
	defer lru.Unlock()
	v, ok := lru.m[key]
	if !ok {
		return nil
	}
	lru.list.Remove(v.node)
	lru.list.AddFirst(v.node)
	for lru.list.length > lru.MaxCap() {
		delete(lru.m, lru.list.tail.val)
		lru.list.RemoveLast()
	}
	return v.val
}

func (lru *defaultLRU) Set(key string, val interface{}) {
	lru.Lock()
	defer lru.Unlock()
	v, ok := lru.m[key]
	if ok {
		lru.list.Remove(v.node)
	}
	node := &Node{
		val: key,
	}
	lru.m[key] = lruVal{
		val:  val,
		node: node,
	}
	lru.list.AddFirst(node)
	for lru.list.length > lru.MaxCap() {
		delete(lru.m, lru.list.tail.val)
		lru.list.RemoveLast()
	}
}

func (lru *defaultLRU) Clear() {
	lru.Lock()
	defer lru.Unlock()
	lru.m = make(map[string]lruVal)
	lru.list = List{}
}

func (lru *defaultLRU) Remove(key string) {
	lru.Lock()
	defer lru.Unlock()
	lru.list.Remove(lru.m[key].node)
	delete(lru.m, key)
}
