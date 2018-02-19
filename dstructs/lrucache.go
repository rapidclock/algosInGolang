package dstructs

import (
	"fmt"
)

type LRUCache struct {
	maxcap int
	curcap int
	vMap map[int]*node
	dNode *dList
}

func NewLRUCache(capacity int) *LRUCache {
	newCache := new(LRUCache)
	newCache.maxcap = capacity
	newCache.vMap = make(map[int]*node)
	newCache.dNode = NewDList()
	return newCache
}

func (this *LRUCache) String() string {
	return fmt.Sprintf("MaxCap : %d, CurCap : %d, map : %v, dList : %v", this.maxcap, len(this.vMap), this.vMap, this.dNode)
}

func (this *LRUCache) Get(key int) int {
	if valNode, ok := this.vMap[key]; ok {
		this.dNode.MoveToHead(valNode)
		return valNode.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key, value int)  {
	if this.maxcap == 0 {
		return
	}
	if valNode, ok := this.vMap[key]; ok {
		valNode.val = value
		this.dNode.MoveToHead(valNode)
	} else {
		if this.curcap < this.maxcap {
			this.curcap += 1
		} else {
			delete(this.vMap, this.dNode.tail.key)
			this.dNode.EvictTail()
		}
		this.dNode.Prepend(key, value)
		this.vMap[key] = this.dNode.head
	}
}