package totp

import (
	"fmt"
	"github.com/hdheid/goutils/fileutil"
	"github.com/hdheid/goutils/otputil"
	"testing"
	"time"
)

func TestTotp_GeneratePwd(t *testing.T) {
	totp := NewTotp(WithSecret("abcd"))
	pwd, err := totp.GeneratePwd()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(pwd)
}

func TestGetQrCode(t *testing.T) {
	totp := NewTotp()
	l := otputil.NewOtpUrl("abcd", "abcd@foxmail.com", totp)
	img, err := l.GetImage(100, 100)
	if err != nil {
		t.Error(err)
	}

	err = fileutil.SaveImageToCurrentDir(img, "qr_code.png")
	if err != nil {
		t.Error(err)
	}

	for {
		pwd, _ := totp.GeneratePwd()
		fmt.Println(pwd)
		time.Sleep(time.Second * 2)
	}
}
