package main

//func main() {
//	mp := treemap.New[int, int](compare.IntMap) // []中分别为键值的数据类型
//
//	// 添加 WithRWMutex() 后，map 为线程安全的 map
//	//mp := treemap.New[int,int](compare.IntMap,treemap.WithRWMutex[int,int]())
//
//	mp.Insert(1, 2) // 1->2，当键存在时，会更新键值
//	mp.Add(2, 3)    // 1->2，2->3 // add 函数就是 insert 函数
//	mp.Add(3, 4)    //  1->2，2->3，3->4
//	mp.Update(1, 3) // 1->3，当没有该键时，update 函数什么也不做
//
//	// 遍历 map
//	for it := mp.Begin(); it.IsValid(); it.Next() {
//		fmt.Println(it.Key(), ",", it.Val()) // 1,3 2,3 3,4
//	}
//
//	// 修改值
//	it := mp.Find(1)
//	it.SetVal(6) // 1->6
//
//	// 判断迭代器是否相等
//	it2 := mp.Begin()
//	fmt.Println(it.Equal(it2)) // true
//}
