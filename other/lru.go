package main

import "container/list"

type entry struct {
	key, value int
}

type LRUCache struct {
	// 容量
	cap int
	// 缓存
	cache map[int]*list.Element
	lst   *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, map[int]*list.Element{}, list.New()}
}

func (c *LRUCache) Get(key int) int {
	e := c.cache[key]
	if e == nil {
		return -1
	}
	// 将缓存移动到链头
	c.lst.MoveToFront(e)
	return e.Value.(entry).value
}

func (c *LRUCache) Put(key, value int) {
	if e := c.cache[key]; e != nil {
		// 如果 key 存在，则修改缓存值
		e.Value = entry{key, value}
		// 将缓存移动到链头
		c.lst.MoveToFront(e)
		return
	}
	// 如果 key 不存在，插入到链头
	c.cache[key] = c.lst.PushFront(entry{key, value})
	if len(c.cache) > c.cap {
		// 如果缓存容量满了，则删除链尾
		delete(c.cache, c.lst.Remove(c.lst.Back()).(entry).key)
	}
}
