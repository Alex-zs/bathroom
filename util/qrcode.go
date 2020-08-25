package util

import (
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

// 创建qrcode，并返回文件路径
func CreateQrcode(content string) string  {
	fileName := "qr.png"
	err := qrcode.WriteFile(content, qrcode.Medium, 256, fileName)
	if err != nil {
		return ""
	}
	dir, _ := os.Getwd()
	return path.Join(dir, fileName)
}

// 获取qrcode的内容
func GetQrcodeContentFromHtml(html string) string {
	r, _ := regexp.Compile("'1\\|(.*)'")
	str := r.FindString(html)
	if len(str) > 0 {
		return str[1:len(str) -1]
	}else {
		return ""
	}
}

// 获取包含登录qrcode的html
func GetHtml() (string, error) {
	resp, err := httpClient.Get("https://vcard.ouc.edu.cn/")
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
