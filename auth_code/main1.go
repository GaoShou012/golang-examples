package main

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func main(){
	id,b64s := GetCaptcha()
	fmt.Println(b64s)
	// 中间可以发给客户端，让他输入然后你来验证
	Verify(id,b64s)
}

//  获取验证码
func GetCaptcha()(string, string){
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)

	// 获取
	id, b64s, err := c.Generate()
	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return "",""
	}
	return id, b64s
}

// 验证验证码
func Verify(id string, val string) bool {
	if id == "" || val == ""{
		return false
	}
	// 同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}