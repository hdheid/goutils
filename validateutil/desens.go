package validateutil

import "strings"

// DesensitizationEmail 邮箱账号脱敏函数 todo：加上邮箱格式是否正确的err
func DesensitizationEmail(email string) string {
	//类似于 123456@qq.com 脱敏成 1****@qq.com
	eList := strings.Split(email, "@") //从@开始分割成几块
	if len(eList) != 2 {               //如果不是两部分，就表示这不是一个正确的邮箱
		return ""
	}
	return eList[0][:1] + "****@" + eList[1]
}

// DesensitizationTel 手机号脱敏函数 todo：加上验证手机号是否正确的err
func DesensitizationTel(tel string) string {
	if len(tel) != 11 { //如果不是是一位数字，就不是一个正确的手机号
		return ""
	}
	return tel[:3] + "****" + tel[7:]
}
