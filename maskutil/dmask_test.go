package maskutil

import (
	"fmt"
	"testing"
)

type Users struct {
	Name     string
	CardID   string
	Email    string `dmask:"eml"`
	TelNum   string `dmask:"tel"`
	Password string `dmask:"pwd"`
}

type Persion struct {
	User   Users
	Length float64
}

func TestDmask(t *testing.T) {
	user := Users{
		Name:     "lh",
		CardID:   "420116200309076611",
		Email:    "422537262@qq.com",
		TelNum:   "18694076387",
		Password: "122po1p11",
	}

	persion := Persion{
		User:   user,
		Length: 175.9,
	}

	maskUser := Dmask(persion)

	fmt.Println(maskUser.User.Email)
	fmt.Println(maskUser.User.TelNum)
	fmt.Println(maskUser.User.Password)
}
