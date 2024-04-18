package menuutil

import (
	"fmt"
	"testing"
)

type Discription struct {
	context string
}

func TestGetMenuList(t *testing.T) {
	menu1 := Menu{id: 1, parentId: -1, name: "一级菜单1", content: Discription{context: "一级菜单1描述"}}
	menu2 := Menu{id: 2, parentId: -1, name: "一级菜单2"}
	menu3 := Menu{id: 3, parentId: 1, name: "二级菜单1"}
	menu4 := Menu{id: 4, parentId: 2, name: "二级菜单2"}
	menu5 := Menu{id: 5, parentId: 3, name: "三级菜单1"}
	menu6 := Menu{id: 6, parentId: 4, name: "三级菜单2"}
	menu7 := Menu{id: 7, parentId: 5, name: "四级菜单2"}
	menu8 := Menu{id: 8, parentId: 6, name: "四级菜单2"}

	list := []Menu{menu1, menu2, menu3, menu4, menu5, menu6, menu7, menu8}

	menuList := GetMenuList(list)

	fmt.Println(menuList)
}
