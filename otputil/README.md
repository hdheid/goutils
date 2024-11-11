<h1 id="QMjCc">OTP</h1>
<h2 id="nYM2v">TOTP</h2>

`totp` 是一个表示基于时间戳算法的一次性密码。基于客户端的动态口令和动态口令验证服务器的时间比对，一般默认每 `30` 秒产生一个新口令，要求客户端和服务器能够十分精确的保持正确相同的时钟，客户端和服务端基于时间计算的动态口令才能一致。

本模块实现了 `totp` 一次性密钥的生成以及对应二维码的生成，通过 `authenticator` 身份验证其扫描二维码后将该动态密码与身份验证器绑定。

<h3 id="TUBU3">代码</h3>

使用代码生成对应二维码：

```go
func main() {
    totp := NewTotp() // 创建一个 totp 实例，可以使用对应 Withxxx 函数配置对应参数
    issuer := "hello"
    account := "1234@email.com"
    l := otputil.NewOtpUrl(issuer, account, totp)
    img, _ := l.GetImage(100, 100) // 生成一个 img 对象
    
    fileutil.SaveImageToCurrentDir(img, "qr_code.png") // 存入二维码到当前目录
}
```

需要获取生成的 `URL` 信息：
```go
func main() {
    totp := NewTotp() // 创建一个 totp 实例，可以使用对应 Withxxx 函数配置对应参数
    issuer := "hello"
    account := "1234@email.com"
    l := otputil.NewOtpUrl(issuer, account, totp)
    
    key, _ := l.GetUrl()
    url := key.Orig
    fmt.Println(url)
}
```

验证输入的一次密码：
```go
func main() {
    t := totp.NewTotp() // 创建一个 totp 实例
    
    var code string
    fmt.Scanf("%s", &code) // 输入一次性密码
    b, _ := t.ValidatePwd(code) // 
    fmt.Println(b)
}
```