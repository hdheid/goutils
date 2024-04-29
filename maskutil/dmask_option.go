package maskutil

import (
	"reflect"
)

func dmaskWithOption(oldVal reflect.Value, option string) reflect.Value {
	switch oldVal.Type().Kind() {
	//case reflect.Struct:
	//	return structDmask(oldVal)
	//case reflect.Float64:
	//	return oldVal
	case reflect.Slice:
		return sliceDmaskWithOption(oldVal, option)
	case reflect.String:
		return StringWithOption(oldVal, option)
	default:
		return oldVal
	}
}

func StringWithOption(oldVal reflect.Value, option string) reflect.Value {
	newVal := reflect.New(oldVal.Type()).Elem()
	newStr := DString(option, oldVal.String())
	newVal.SetString(newStr)
	return newVal
}

func sliceDmaskWithOption(oldVal reflect.Value, tagCtx string) reflect.Value {
	t := oldVal.Type()
	newVal := reflect.MakeSlice(t, oldVal.Len(), oldVal.Len())

	for i := 0; i < oldVal.Len(); i++ {
		val := oldVal.Index(i)
		switch t.Elem().Kind() { // 不加elem，获取的是切片类型，加上elem，获取的是切片的每一个元素的类型，使用 val.Kind() 能达到同样的效果
		case reflect.String:
			s := DString(tagCtx, val.String())
			newVal.Index(i).SetString(s)
		default:
			return oldVal
			//subField := dmask(val, tagCtx)
			//newVal.Index(i).Set(subField)
		}
	}

	return newVal
}
