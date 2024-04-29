package otherutil

import (
	"errors"
	"reflect"
	"strings"
)

// Contains 判断某个元素是否存在于对应 carrier 的中
// todo: 待完善，后续支持更多数据类型
func Contains(obj, carrier interface{}) (bool, error) {
	carrierType := reflect.ValueOf(carrier) //判断载体类型
	objType := reflect.ValueOf(obj)
	if carrierType.Kind() != objType.Kind() {
		return false, errors.New("Type not match ")
	}

	switch carrierType.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < carrierType.Len(); i++ {
			if carrierType.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if carrierType.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	case reflect.String:
		objStr := obj.(string)
		carrierStr := carrier.(string)
		return strings.Contains(carrierStr, objStr), nil
	default:
	}

	return false, nil
}
