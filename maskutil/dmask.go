package maskutil

import (
	"github.com/hdheid/goutils/common"
	"github.com/hdheid/goutils/validateutil"
	"reflect"
	"regexp"
)

//// Dmask 将结构体中需要脱敏的部分进行脱敏，传地址
//func Dmask[T any](obj T) T {
//	return dmask(reflect.ValueOf(obj), reflect.Value{}).Interface().(T)
//
//	//switch v.Type().Kind() {
//	//case reflect.Struct:
//	//	NewStruct := dmask(v, reflect.Value{}).Interface()
//	//	return NewStruct.(T)
//	//default:
//	//	panic("unhandled default case")
//	//}
//}
//
//func dmask(oldVal reflect.Value, newVal reflect.Value) reflect.Value {
//	switch oldVal.Type().Kind() {
//	case reflect.Struct:
//		return structDmask(oldVal, newVal)
//	case reflect.Float64:
//		return oldVal
//	default:
//		panic("unhandled default case")
//	}
//}
//
//func structDmask(oldVal, newVal reflect.Value) reflect.Value {
//	t := oldVal.Type()
//	if !newVal.IsValid() { // 赋初始值
//		newVal = reflect.New(t).Elem()
//	}
//
//	for i := 0; i < t.NumField(); i++ {
//		field := t.Field(i)
//
//		tagCtx := field.Tag.Get(common.Tag)
//
//		//判断字段的类型，是否是字符串
//		switch field.Type.Kind() {
//		case reflect.String:
//			s := StructString(tagCtx, oldVal.Field(i).String())
//			newVal.Field(i).SetString(s)
//
//		default:
//			subField := dmask(oldVal.Field(i), newVal.Field(i))
//			newVal.Field(i).Set(subField)
//		}
//	}
//
//	return newVal
//}

// Dmask 将结构体中需要脱敏的部分进行脱敏，传地址
func Dmask[T any](obj T) T {
	return dmask(reflect.ValueOf(obj)).Interface().(T)
}

func dmask(oldVal reflect.Value) reflect.Value {
	switch oldVal.Type().Kind() {
	case reflect.Struct:
		return structDmask(oldVal)
	case reflect.Float64:
		return oldVal
	default:
		panic("unhandled default case")
	}
}

func structDmask(oldVal reflect.Value) reflect.Value {
	t := oldVal.Type()
	newVal := reflect.New(t).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tagCtx := field.Tag.Get(common.Tag)

		//判断字段的类型，是否是字符串
		switch field.Type.Kind() {
		case reflect.String:
			s := StructString(tagCtx, oldVal.Field(i).String())
			newVal.Field(i).SetString(s)

		default:
			subField := dmask(oldVal.Field(i))
			newVal.Field(i).Set(subField)
		}
	}

	return newVal
}

func StructString(tagCtx, fieldCtx string) string {
	switch tagCtx {
	case common.Email:
		return EmlDmask(fieldCtx)
	case common.Tel:
		return TelDmask(fieldCtx)
	case common.Password:
		pwd, _ := PwdDmasl(fieldCtx)
		return pwd
	default:
		return fieldCtx
	}
}

func EmlDmask(eml string) string {
	re := regexp.MustCompile(`(\d)(\d*)(@\w+\.\w+)`)    // 匹配邮箱中的数字部分
	hiddenEmail := re.ReplaceAllString(eml, "$1****$3") // 替换数字部分为*
	return hiddenEmail
}

func TelDmask(tel string) string {
	re := regexp.MustCompile(`(\d{3})(\d{4})(\d{4})`)         // 匹配电话号码中的前三位、中间四位和最后四位
	hiddenPhoneNumber := re.ReplaceAllString(tel, "$1****$3") // 替换中间四位为*
	return hiddenPhoneNumber
}

// PwdDmasl 针对密码，使用加密算法
func PwdDmasl(pwd string) (string, error) {
	hashPwd, err := validateutil.PasswordHash(pwd)
	if err != nil {
		return "", err
	}
	return hashPwd, nil
}

// todo：密码加密，map，切片，普通字符串等等
