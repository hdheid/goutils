package skip_list

import (
	"github.com/hdheid/goutils/common/compare"
	"sync"
	"testing"
)

func TestInsert(t *testing.T) {
	skipList := New[int, int](compare.IntSkipList)

	mp := make(map[int]int)
	for i := 0; i < 100; i++ {
		key := i * 50
		skipList.Insert(key, i)
		//mp[key] = i
	}

	for i := 0; i < 100; i++ {
		mp[i*50] = i
	}

	lists := skipList.Traversal()
	mp2 := make(map[int]int)
	for _, list := range lists {
		mp2[list.key] = list.val
	}

	for k, v := range mp {
		if v != mp2[k] {
			t.Errorf("err")
		}
	}

	for k, v := range mp {
		val, _ := skipList.Find(k)
		if v != val {
			t.Error("skipList error! ", k, " ", v, " ", val)
		}
	}
}

func TestConcurrency(t *testing.T) {
	skipList := New[int, int](func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else if a == b {
			return 0
		}
		return 0
	}, WithRWMutex[int, int]())

	mp := make(map[int]int)

	wg := sync.WaitGroup{}
	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		i := i
		go func() {
			defer wg.Done()
			key := i * 50
			skipList.Insert(key, i)
		}()
	}

	wg.Wait()

	for i := 0; i < 10000; i++ {
		mp[i*50] = i
	}

	lists := skipList.Traversal()
	mp2 := make(map[int]int)
	for _, list := range lists {
		mp2[list.key] = list.val
	}

	for k, v := range mp {
		if v != mp2[k] {
			t.Errorf("err")
		}
	}

	for k, v := range mp {
		val, _ := skipList.Find(k)
		if v != val {
			t.Error("skipList error! ", k, " ", v, " ", val)
		}
	}
}

func TestRemove(t *testing.T) {
	var count = 10000

	skipList := New[int, int](func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else if a == b {
			return 0
		}
		return 0
	})

	mp := make(map[int]int)
	for i := 0; i < count; i++ {
		key := i * 50
		skipList.Insert(key, i)
		//mp[key] = i
	}

	for i := 0; i < count; i++ {
		mp[i*50] = i
	}

	lists := skipList.Traversal()
	mp2 := make(map[int]int)
	for _, list := range lists {
		mp2[list.key] = list.val
	}

	for i := 0; i < count-645; i++ {
		key := i * 50
		skipList.Remove(key)
		delete(mp, i*50)
	}

	for k, v := range mp {
		val, _ := skipList.Find(k)
		if v != val {
			t.Error("skipList error! ", k, " ", v, " ", val)
		}
	}
}
