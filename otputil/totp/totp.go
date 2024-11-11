package totp

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/hdheid/goutils/common"
	"github.com/hdheid/goutils/validateutil"
	"hash"
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 一次性密码校验，部分参考自：https://github.com/pquerna/otp

// 密码长度、密码有效期（时间间隔）、加密算法、密钥
const (
	SHA1 Algorithm = iota
	SHA256
	SHA512
	MD5
)

type Algorithm int

type OpFunc func(t *Totp)

type Totp struct {
	Secret        string    // 密钥，用于生成密码
	Expiration    int       // 密码有效期
	Size          int       // 密码长度
	HashAlgorithm Algorithm // 加密算法
}

// WithAlgorithm 设置使用的，默认SHA1，可选择有 “SHA1，SHA256，SHA512，MD5”，为保证兼容性，建议使用默认哈希算法
func WithAlgorithm(algorithm Algorithm) OpFunc {
	return func(t *Totp) {
		t.HashAlgorithm = algorithm
	}
}

// WithSize 设置生成一次性密码长度
func WithSize(size int) OpFunc {
	return func(t *Totp) {
		t.Size = size
	}
}

// WithExpiration 设置密码刷新间隔时间
func WithExpiration(expiration int) OpFunc {
	return func(t *Totp) {
		t.Expiration = expiration
	}
}

// WithSecret 设置一次性密码密钥，应符合 base32 字符集
func WithSecret(secret string) OpFunc {
	return func(t *Totp) {
		t.Secret = secret
	}
}

func NewTotp(ops ...OpFunc) *Totp {
	t := &Totp{ // 默认值
		Secret:        validateutil.GenValidateString(common.Default_Secret_Size, 5),
		Expiration:    common.Default_Expiration, // 默认30秒
		Size:          common.Default_Size,       // 默认6
		HashAlgorithm: SHA1,                      // 默认sha1
	}

	for _, op := range ops {
		op(t)
	}

	return t
}

// GeneratePwd 生成密码
func (t *Totp) GeneratePwd() (pwd string, err error) {
	// 第一部，获取密钥，base32编码必须是8的倍数，全部大写，如果不是，需要填充’=‘
	secret := sercetFormat(t.Secret)
	sBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	// 第二步，获取当前时间戳
	c := uint64(time.Now().Unix()) / uint64(t.Expiration)
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, c)

	// 获取密码
	mac := hmac.New(t.HashAlgorithm.GetAlgorithm(), sBytes)
	mac.Write(buf)
	sum := mac.Sum(nil)

	offset := sum[len(sum)-1] & 0xf // 获取到偏移量，为最后一位低四位
	// 从偏移位开始取四个字节作为值
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))
	// 得到取模后的数字
	mod := int32(value % int64(math.Pow10(t.Size)))

	// 位数不足前面取零
	f := fmt.Sprintf("%%0%dd", t.Size)
	password := fmt.Sprintf(f, mod)

	return password, nil
}

func (t *Totp) ValidatePwd(pwd string) (bool, error) {
	pwd = strings.TrimSpace(pwd)
	if len(pwd) != t.Size {
		return false, errors.New("长度不一致")
	}

	newPwd, err := t.GeneratePwd()
	if err != nil {
		return false, err
	}

	// 涨见识了
	if subtle.ConstantTimeCompare([]byte(pwd), []byte(newPwd)) == 1 {
		return true, nil
	}

	return false, errors.New("认证失败")
}

func (t *Totp) GetUrl(issuer, account string) url.URL {
	u := url.Values{}
	u.Set(common.SECRET, t.Secret)
	u.Set(common.ISSUER, issuer)
	u.Set(common.PERIOD, strconv.FormatUint(uint64(t.Expiration), 10))
	u.Set(common.ALGORITHM, t.HashAlgorithm.String())
	u.Set(common.DIGITS, strconv.Itoa(t.Size))

	return url.URL{
		Scheme:   "otpauth",
		Host:     "totp",
		Path:     "/" + issuer + ":" + account,
		RawQuery: EncodeQuery(u),
	}
}

func (a Algorithm) GetAlgorithm() (h func() hash.Hash) {
	switch a {
	case SHA1:
		return sha1.New
	case SHA256:
		return sha256.New
	case SHA512:
		return sha512.New
	case MD5:
		return md5.New
	}

	panic("err: Incompatible algorithms")
}

func (a Algorithm) String() string {
	switch a {
	case SHA1:
		return "SHA1"
	case SHA256:
		return "SHA256"
	case SHA512:
		return "SHA512"
	case MD5:
		return "MD5"
	}

	panic("err: Incompatible algorithms")
}

func sercetFormat(pwd string) string {
	pwd = strings.TrimSpace(pwd)
	if n := len(pwd) % 8; n != 0 { // 补充=保证为8的倍数
		pwd = pwd + strings.Repeat("=", 8-n)
	}

	return strings.ToUpper(pwd)
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
