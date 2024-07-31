https://github.com/emirpasic/gods good data struct by go

# 数据结构
- [List](#list)
- [Deque](#deque)
- [Map](#map)


## List

## Deque

## Map
Map 是一个基于红黑树实现的有序关联容器，平均插入、删除、查找时间复杂度为 O(logn)。Map 可以自选为线程安全模式。下面是 map 的使用：
```go
package main

import (
	"github.com/hdheid/goutils/common/compare"
	treemap "github.com/hdheid/goutils/structuitl/data_struct/map"
)

func main() {
	mp := treemap.New[int, int](compare.IntMap) // []中分别为键值的数据类型

	// 添加 WithRWMutex() 后，map 为线程安全的 map
	//mp := treemap.New[int,int](compare.IntMap,treemap.WithRWMutex[int,int]())

	mp.Insert(1, 2) // 1->2，当键存在时，会更新键值
	mp.Add(2, 3)    // 1->2，2->3 // add 函数就是 insert 函数
	mp.Add(3, 4)    //  1->2，2->3，3->4
	mp.Update(1, 3) // 1->3，当没有该键时，update 函数什么也不做

	val, ok := mp.Get(2) // 3,true，Get 函数会查询键对应的值，如果不存在，ok 会返回 false

	mp.Delete(1) // 2->3，3->4，delete 函数删除对应键值对

	it1 := mp.Find(3) // 返回指向键值对 3->4 的迭代器
	_ := mp.Begin()   // 返回指向 map 第一个节点的迭代器
	mp.Erase(it1)     // 2->3，erase 函数删除迭代器所指向的键值对
	
	mp.Clear() // 清空 map
	mp.Size()  // 0
}
```

map 中存在迭代器，可以通过迭代器访问 map：
```go
package main

import (
	"fmt"
	"github.com/hdheid/goutils/common/compare"
	treemap "github.com/hdheid/goutils/structuitl/data_struct/map"
)

func main() {
	mp := treemap.New[int, int](compare.IntMap) // []中分别为键值的数据类型

	// 添加 WithRWMutex() 后，map 为线程安全的 map
	//mp := treemap.New[int,int](compare.IntMap,treemap.WithRWMutex[int,int]())

	mp.Insert(1, 2) // 1->2，当键存在时，会更新键值
	mp.Add(2, 3)    // 1->2，2->3 // add 函数就是 insert 函数
	mp.Add(3, 4)    //  1->2，2->3，3->4
	mp.Update(1, 3) // 1->3，当没有该键时，update 函数什么也不做

	// 遍历 map
	for it := mp.Begin(); it.IsValid(); it.Next() {
		fmt.Println(it.Key(), ",", it.Val()) // 1,3 2,3 3,4
	}

	// 修改值
	it := mp.Find(1)
	it.SetVal(6) // 1->6

	// 判断迭代器是否相等
	it2 := mp.Begin()
	fmt.Println(it.Equal(it2)) // true

	cloneIt := it.Clone() // 拷贝一个新的迭代器（深拷贝）
}
```

# 计划
### 线程安全
- [x] bloomFilter
- [x] queue
- [x] stack
- [x] priorityQueue
- [x] skipList
- [x] map
- [ ] set
- [ ] multimap
- [ ] multiset
- [ ] ...




### 非线程安全
- [x] bitMap
- [x] deque
- [x] heap
- [x] rbTree
- [ ] pair
- [ ] list
- [ ] ...


[回到顶部](#数据结构)