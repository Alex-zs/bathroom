package util

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
)

func WsClient(qrcode string) (string, error){
	// 建立连接
	u := url.URL{Scheme: "wss", Host: "vcard.ouc.edu.cn:20087"}
	dialer := websocket.Dialer{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return "", err
	}

	defer c.Close()
	c.WriteMessage(websocket.TextMessage, []byte("(" + qrcode + ")"))

	fmt.Printf("请使用海大e卡通扫描二维码...\n")

	_, message, err := c.ReadMessage()
	fmt.Printf("登录成功.\n")
	return string(message), nil
}
