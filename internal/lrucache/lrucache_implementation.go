package lrucache

import (
	"container/list"
)

type LRUEntity struct {
	capacity int
	mapCache map[string]*list.Element
	list     *list.List
}

type data struct {
	k, v string
}

func NewLRUCache(n int) LRUCache {
	return &LRUEntity{
		capacity: n,
		mapCache: map[string]*list.Element{},
		list:     list.New(),
	}
}

func (l *LRUEntity) Add(key, value string) bool {
	if val, ok := l.mapCache[key]; ok {
		l.list.MoveToFront(val)
		val.Value = data{k: key, v: value} // перезапись значения, вдруг новое по ключу
		return false
	}

	if l.list.Len() == l.capacity {
		l.Remove(l.list.Back().Value.(data).k)
	}

	element := l.list.PushFront(data{k: key, v: value})
	l.mapCache[key] = element

	//fmt.Println(l.mapCache)
	//fmt.Println(l.list.Len())

	return true
}

func (l *LRUEntity) Get(key string) (value string, ok bool) {
	ok = false
	if val, okMap := l.mapCache[key]; okMap {
		l.list.MoveToFront(val)
		value = val.Value.(data).v
		ok = true
	}
	return
}

func (l *LRUEntity) Remove(key string) (ok bool) {
	ok = false
	if val, okMap := l.mapCache[key]; okMap {
		delete(l.mapCache, key)
		l.list.Remove(val)
		ok = true
	}
	//fmt.Println(l.mapCache)
	//fmt.Println(l.list.Len())

	return
}
