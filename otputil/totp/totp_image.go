package totp

import (
	"errors"
	"github.com/hdheid/goutils/common"
	"github.com/hdheid/goutils/otputil/qr_code"
	"image"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type OtpUrl struct {
	Issuer  string // 公司或者组织的名字
	Account string // eg，email的地址
	totp    *Totp  // 基于时间的一次性密码
}

// Key represents an TOTP or HTOP key.
type Key struct {
	orig string
	url  *url.URL
}

func NewOtpUrl(issuer string, account string, totp *Totp) *OtpUrl {
	return &OtpUrl{
		Issuer:  issuer,
		Account: account,
		totp:    totp,
	}
}

func (o *OtpUrl) GetUrl() (*Key, error) {
	if o.Issuer == "" {
		return nil, errors.New("issuer 不能为空")
	}

	if o.Account == "" {
		return nil, errors.New("account 不能为空")
	}

	u := url.Values{}
	u.Set(common.SECRET, o.totp.secret)
	u.Set(common.ISSUER, o.Issuer)
	u.Set(common.PERIOD, strconv.FormatUint(uint64(o.totp.expiration), 10))
	u.Set(common.ALGORITHM, o.totp.hashAlgorithm.String())
	u.Set(common.DIGITS, strconv.Itoa(o.totp.size))

	otpUrl := url.URL{
		Scheme:   "otpauth",
		Host:     "totp",
		Path:     "/" + o.Issuer + ":" + o.Account,
		RawQuery: EncodeQuery(u)}

	return &Key{
		orig: otpUrl.String(),
		url:  &otpUrl,
	}, nil

	// otpauth://totp/Example:alice@google.com?secret=JBSWY3DPEHPK3PXP&issuer=Example
}

func (k *Key) Image(width int, height int) (image.Image, error) {
	b, err := qr_code.Encode(k.orig, qr_code.M, qr_code.Auto)
	if err != nil {
		return nil, err
	}

	b, err = qr_code.Scale(b, width, height)

	if err != nil {
		return nil, err
	}

	return b, nil
}

// EncodeQuery is a copy-paste of url.Values.Encode, except it uses %20 instead
// of + to encode spaces. This is necessary to correctly render spaces in some
// authenticator apps, like Google Authenticator.
func EncodeQuery(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := url.PathEscape(k) // changed from url.QueryEscape
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.PathEscape(v)) // changed from url.QueryEscape
		}
	}
	return buf.String()
}
