package menuutil

import "github.com/hdheid/goutils/common"

/*
任何需要实现多级列表的功能，除了固定的几个字段外，其他字段都可以单独定义一个结构体，存在 content 字段中
*/

type Menu struct {
	id       int64
	parentId int64
	name     string
	content  interface{}
	subList  []*Menu
}

func GetMenuList(list []Menu) (menuList []Menu) {
	menuMap := make(map[int64]*Menu)
	for i := range list {
		menuMap[list[i].id] = &list[i]
	}

	result := make([]Menu, 0, len(list))
	for i := range list {
		if list[i].parentId == -1 {
			continue // 这里没有continue导致下面一段 parentMenu 的值为nil，因此 nil.subList 导致空指针解引用错误
		}
		parentMenu, ok := menuMap[list[i].parentId] // 这里需要获取到地址，否则将是深拷贝，下面 append 的数据与 map 中的 menu 无关
		if !ok {
			continue
		}

		if parentMenu.subList == nil {
			parentMenu.subList = []*Menu{}
		}
		parentMenu.subList = append(parentMenu.subList, &list[i])
	}

	for _, menu := range list {
		if menu.parentId == common.ROOTDIR {
			result = append(result, menu)
		}
	}

	return result
}
