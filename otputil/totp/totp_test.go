package totp

import (
	"fmt"
	"github.com/hdheid/goutils/fileutil"
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
	l := NewOtpUrl("abcd", "abcd@foxmail.com", totp)
	k, _ := l.GetUrl()

	img, err := k.Image(100, 100)
	if err != nil {
		t.Error(err)
	}

	err = fileutil.SaveImageToCurrentDir(img, "qr_code.png")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(k.orig)

	for {
		pwd, _ := totp.GeneratePwd()
		fmt.Println(pwd)
		time.Sleep(time.Second * 2)
	}
}
