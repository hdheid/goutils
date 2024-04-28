package validateutil

import "testing"

func TestPasswordHashAndVerify(t *testing.T) {
	pwd := "1234567"
	hashPwd, err := PasswordHash(pwd)
	if err != nil {
		t.Error(err)
	}

	ok := PasswordVerify(hashPwd, pwd)
	if !ok {
		t.Errorf("error")
	}
}
