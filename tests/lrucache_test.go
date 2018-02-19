package tests

import (
	ds "../dstructs"
	"testing"
)

func TestLRUZeroCapacity(t *testing.T) {
	lruCache := ds.NewLRUCache(0)
	if lruCache.Get(1) != -1 {
		t.Errorf("Zero Capacity Test : get(1) should be -1 but was not.")
	}
	lruCache.Put(1, 1)
	if lruCache.Get(1) != -1 {
		t.Errorf("Zero Capacity Test : get(1) should be -1 but was not.")
	}
	lruCache.Put(2, 2)
	if lruCache.Get(1) != -1 {
		t.Errorf("Zero Capacity Test : get(1) should be -1 but was not.")
	}
	if lruCache.Get(2) != -1 {
		t.Errorf("Zero Capacity Test : get(2) should be -1 but was not.")
	}
}

func TestLRUCapacityOne(t *testing.T) {
	lruCache := ds.NewLRUCache(1)
	if lruCache.Get(1) != -1 {
		t.Errorf("One Capacity Test : get from empty cache should be -1 but wasn't")
	}
	lruCache.Put(1, 1)
	if lruCache.Get(1) != 1 {
		t.Errorf("One Capacity Test : get(1) should be 1 but was not.")
	}
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	if lruCache.Get(1) != -1 {
		t.Errorf("One Capacity Test : get(1) should be -1 but was not.")
	}
	if lruCache.Get(2) != -1 {
		t.Errorf("One Capacity Test : get(2) should be -1 but was not.")
	}
	if lruCache.Get(3) != 3 {
		t.Errorf("One Capacity Test : get(3) should be 3 but was not.")
	}
}

func TestLRUCapacityOneReplaceValue(t *testing.T) {
	lruCache := ds.NewLRUCache(2)
	lruCache.Put(1, 1)
	if lruCache.Get(1) != 1 {
		t.Errorf("Replace Value Test : get(1) should be 1 but was not.")
	}
	lruCache.Put(1, 2)
	if val := lruCache.Get(1); val != 2 {
		t.Errorf("Replace Value Test : get(1) : Expected 2, Actual %d", val)
	}
	lruCache.Put(3, 3)
	if val := lruCache.Get(1); val != 2 {
		t.Errorf("Replace Value Test : get(1) : Expected 2, Actual %d", val)
	}
	if val := lruCache.Get(3); val != 3 {
		t.Errorf("Replace Value Test : get(3) : Expected 3, Actual %d", val)
	}
	lruCache.Put(3, 4)
	if val := lruCache.Get(1); val != 2 {
		t.Errorf("Replace Value Test : get(1) : Expected 2, Actual %d", val)
	}
	if val := lruCache.Get(3); val != 4 {
		t.Errorf("Replace Value Test : get(3) : Expected 4, Actual %d", val)
	}
}

func TestLRUNormalCapacity(t *testing.T) {
	lruCache := ds.NewLRUCache(3)
	lruCache.Put(1,1)
	lruCache.Put(2,2)
	lruCache.Put(3,3)
	if val := lruCache.Get(3); val != 3 {
		t.Errorf("Normal Test : get(3) : Expected 3, Actual %d", val)
	}
	if val := lruCache.Get(2); val != 2 {
		t.Errorf("Normal Test : get(2) : Expected 2, Actual %d", val)
	}
	if val := lruCache.Get(1); val != 1 {
		t.Errorf("Normal Test : get(1) : Expected 1, Actual %d", val)
	}
	lruCache.Put(3,4)
	if val := lruCache.Get(3); val != 4 {
		t.Errorf("Normal Test : get(3) : Expected 4, Actual %d", val)
	}
	if val := lruCache.Get(2); val != 2 {
		t.Errorf("Normal Test : get(2) : Expected 2, Actual %d", val)
	}
	if val := lruCache.Get(1); val != 1 {
		t.Errorf("Normal Test : get(1) : Expected 1, Actual %d", val)
	}
	lruCache.Put(5,5)
	if val := lruCache.Get(2); val != 2 {
		t.Errorf("Normal Test : get(2) : Expected 2, Actual %d", val)
	}
	if val := lruCache.Get(3); val != -1 {
		t.Errorf("Normal Test : get(3) : Expected -1, Actual %d", val)
	}
	lruCache.Put(6,6)
	if val := lruCache.Get(3); val != -1 {
		t.Errorf("Normal Test : get(3) : Expected -1, Actual %d", val)
	}
	if val := lruCache.Get(1); val != -1 {
		t.Errorf("Normal Test : get(1) : Expected -1, Actual %d\nCache : %v", val, lruCache)
	}
	lruCache.Put(7,7)
	if val := lruCache.Get(2); val != 2 {
		t.Errorf("Normal Test : get(2) : Expected 2, Actual %d", val)
	}
	if val := lruCache.Get(6); val != 6 {
		t.Errorf("Normal Test : get(6) : Expected 6, Actual %d", val)
	}
	if val := lruCache.Get(7); val != 7 {
		t.Errorf("Normal Test : get(7) : Expected 7, Actual %d", val)
	}
	if val := lruCache.Get(5); val != -1 {
		t.Errorf("Normal Test : get(5) : Expected -1, Actual %d\nCache : %v", val, lruCache)
	}
}

func TestTenCapacityCache(t *testing.T) {
	cache := ds.NewLRUCache(10)
	cache.Put(10,13)
	cache.Put(3,17)
	cache.Put(6,11)
	cache.Put(10,5)
	cache.Put(9,10)
	cache.Get(13)
	cache.Put(2,19)
	cache.Get(2)
	cache.Get(3)
	cache.Put(5,25)
	cache.Get(8)
	cache.Put(9,22)
	cache.Put(5,5)
	cache.Put(1,30)
	cache.Get(11)
	cache.Put(9,12)
	cache.Get(7)
	cache.Get(5)
	cache.Get(8)
	cache.Get(9)
	cache.Put(4,30)
	cache.Put(9,3)
	if val := cache.Get(9); val != 3 {
		t.Errorf("10 Cap Test : get(9) : Expected 3, Actual %d\nCache : %v", val, cache)
	}
	if val := cache.Get(10); val != 5 {
		t.Errorf("10 Cap Test : get(10) : Expected 5, Actual %d\nCache : %v", val, cache)
	}
	cache.Get(10)
	cache.Put(6,14)
	cache.Put(3,1)
	cache.Get(3)
	cache.Put(10,11)
	cache.Get(8)
	cache.Put(2,14)
	cache.Get(1)
	cache.Get(5)
	cache.Get(4)
	cache.Put(11,4)
	cache.Put(12,24)
	cache.Put(5,18)
	if val := cache.Get(13); val != -1 {
		t.Errorf("10 Cap Test : get(13) : Expected -1, Actual %d\nCache : %v", val, cache)
	}
	cache.Put(7,23)
	cache.Get(8)
	cache.Get(12)
	cache.Put(3,27)
	cache.Put(2,12)
	if val := cache.Get(5); val != 18 {
		t.Errorf("10 Cap Test : get(5) : Expected 18, Actual %d\nCache : %v", val, cache)
	}
	cache.Put(2,9)
	cache.Put(13,4)
	cache.Put(8,18)
	cache.Put(1,7)
	if val := cache.Get(6); val != -1 {
		t.Errorf("10 Cap Test : get(6) : Expected -1, Actual %d\nCache : %v", val, cache)
	}
	cache.Put(9,29)
	cache.Put(8,21)
	cache.Get(5)
	cache.Put(6,30)
	cache.Put(1,12)
	cache.Get(10)
	cache.Put(4,15)
	cache.Put(7,22)
	cache.Put(11,26)
	cache.Put(8,17)
	cache.Put(9,29)
	if val := cache.Get(5); val != 18 {
		t.Errorf("10 Cap Test : get(5) : Expected 18, Actual %d\nCache : %v", val, cache)
	}
	cache.Put(3,4)
	if val := cache.Get(13); val != 4 {
		t.Errorf("10 Cap Test : get(13) : Expected 4, Actual %d\nCache : %v", val, cache)
	}
}