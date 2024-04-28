package maskutil

import (
	"github.com/hdheid/goutils/common"
	"github.com/hdheid/goutils/validateutil"
	"reflect"
	"regexp"
)

// Dmask 将结构体中需要脱敏的部分进行脱敏，传地址
func Dmask[T any](obj T) T {
	v := reflect.ValueOf(obj)

	switch v.Type().Kind() {
	case reflect.Struct:
		NewStruct := structDmask(v).Interface()
		return NewStruct.(T)
	default:
		panic("unhandled default case")
	}

	return obj
}

func structDmask(v reflect.Value) reflect.Value {
	t := v.Type()
	newStruct := reflect.New(t).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tagCtx := field.Tag.Get(common.Tag)

		//判断字段的类型，是否是字符串
		switch field.Type.Kind() {
		case reflect.String:
			s := StructString(tagCtx, v.Field(i).String())
			newStruct.Field(i).SetString(s)
		default:
			panic("unhandled default case")
		}
	}

	return newStruct
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
