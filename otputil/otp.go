package otputil

import "net/url"

type Otp interface {
	GeneratePwd() (pwd string, err error)
	ValidatePwd(pwd string) (bool, error)
	GetUrl(issuer, account string) url.URL
}
