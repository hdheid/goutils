package sliceutil

import (
	"fmt"
	"testing"
)

type Map struct {
	CountryName string
	Count       int
}

func TestMergeSlice(t *testing.T) {
	var listA = []Map{
		{
			CountryName: "中国",
			Count:       1,
		},
		{
			CountryName: "美国",
			Count:       2,
		},
		{
			CountryName: "英国",
			Count:       2,
		},
	}

	var listB = []Map{
		{
			CountryName: "中国",
			Count:       1,
		},
		{
			CountryName: "美国",
			Count:       2,
		},
	}

	var expect = []Map{
		{
			CountryName: "中国",
			Count:       2,
		},
		{
			CountryName: "美国",
			Count:       4,
		},
		{
			CountryName: "英国",
			Count:       2,
		},
	}

	mergeList := MergeSliceWithAny[Map](listA, listB, func(a, b Map) (Map, bool) {
		if a.CountryName == b.CountryName {
			return Map{
				CountryName: a.CountryName,
				Count:       a.Count + b.Count,
			}, true
		}

		return Map{}, false
	})

	for i := range mergeList {
		if mergeList[i].Count != expect[i].Count || mergeList[i].CountryName != expect[i].CountryName {
			t.Errorf("expect %v, got %v", expect[i], mergeList[i])
		}
	}
}

func TestDeleteIndSliceWithAnyTest(t *testing.T) {
	age := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//ecpectAge := []int{2,3,4,5,6,7,8,9}

	age = DeleteMultiIndSliceWithAny(age, 7, 8, 1, 2)

	fmt.Println(age)
}
