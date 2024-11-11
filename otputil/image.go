package otputil

import (
	"errors"
	"github.com/hdheid/goutils/otputil/qr_code"
	"image"
	"net/url"
)

type OtpUrl struct {
	Issuer  string // 公司或者组织的名字
	Account string // eg，email的地址
	Otp     Otp    // 一次性密码
}

// Key represents an TOTP or HTOP key.
type Key struct {
	Orig string
	Url  *url.URL
}

func NewOtpUrl(issuer string, account string, otp Otp) *OtpUrl {
	return &OtpUrl{
		Issuer:  issuer,
		Account: account,
		Otp:     otp,
	}
}

func (o *OtpUrl) GetUrl() (*Key, error) {
	if o.Issuer == "" {
		return nil, errors.New("issuer 不能为空")
	}

	if o.Account == "" {
		return nil, errors.New("account 不能为空")
	}

	if o.Otp == nil {
		panic("otp 不能为空")
	}
	otpUrl := o.Otp.GetUrl(o.Issuer, o.Account)

	return &Key{
		Orig: otpUrl.String(),
		Url:  &otpUrl,
	}, nil

	// otpauth://totp/Example:alice@google.com?secret=JBSWY3DPEHPK3PXP&issuer=Example
}

func (k *Key) Image(width int, height int) (image.Image, error) {
	b, err := qr_code.Encode(k.Orig, qr_code.M, qr_code.Auto)
	if err != nil {
		return nil, err
	}

	b, err = qr_code.Scale(b, width, height)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// GetImage 对上面两个函数的封装
func (o *OtpUrl) GetImage(width int, height int) (img image.Image, err error) {
	key, err := o.GetUrl()
	if err != nil {
		return nil, err
	}

	img, err = key.Image(width, height)
	if err != nil {
		return nil, err
	}

	return img, nil
}
