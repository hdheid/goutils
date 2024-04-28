package validateutil

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash 密码加密，默认 cost 为 10
func PasswordHash(pwd string, costs ...int) (string, error) {
	cost := bcrypt.DefaultCost

	if len(costs) > 1 {
		return "", errors.New("cost only need one")
	}
	if len(costs) == 1 {
		cost = costs[0]
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), cost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func PasswordVerify(pwdHash, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(pwd))
	return err == nil
}
