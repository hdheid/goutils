package validateutil

import "testing"

func TestDesensitizationEmail(t *testing.T) {
	email := "123456789@qq.com"
	expectation := "1****@qq.com"

	desensEmail := DesensitizationEmail(email)

	if desensEmail != expectation {
		t.Errorf("DesensitizationEmail(%s) = %s, want %s", email, desensEmail, expectation)
	}
}

func TestDesensitizationTel(t *testing.T) {
	tel := "12345678923"
	expectation := "123****8923"

	desensTel := DesensitizationTel(tel)

	if desensTel != expectation {
		t.Errorf("DesensitizationEmail(%s) = %s, want %s", tel, desensTel, expectation)
	}
}
