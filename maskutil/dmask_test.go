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

func TestDmask(t *testing.T) {
	user := Users{
		Name:     "lh",
		CardID:   "420116200309076611",
		Email:    "422537262@qq.com",
		TelNum:   "18694076387",
		Password: "122po1p11",
	}

	maskUser := Dmask(user)

	fmt.Println(maskUser.Email)
	fmt.Println(maskUser.TelNum)
	fmt.Println(maskUser.Password)
}
