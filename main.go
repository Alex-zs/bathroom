package main

import (
	"fmt"
	"github.com/Alex-zs/bathroom/util"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// 打印banner
	fmt.Print(util.Banner)

	// 获取qrcode
	html, err := util.GetHtml()
	if err != nil {
		fmt.Printf("%s %s\n", "请求失败，请检查网络环境", err)
		return
	}
	qrcodeContent := util.GetQrcodeContentFromHtml(html)

	// 创建验证码图片并打开
	qrcodeFilePath := util.CreateQrcode(qrcodeContent)
	fileURL := "file:///" + qrcodeFilePath
	goos := runtime.GOOS
	switch goos {
	case "darwin":
		exec.Command(`open`, fileURL).Start()
	case "windows":
		exec.Command(`cmd`, `/c`, `start`, fileURL).Start()
	case "linux":
		exec.Command(`xdg-open`, fileURL).Start()
	}
	// 连接websocket并获取令牌
	account, err := util.WsClient(qrcodeContent[2:])
	if err != nil {
		fmt.Printf("连接失败，请检查网络。原因:%s\n", err)
		return
	}

	// 删除qrcode文件
	_ = os.Remove(qrcodeFilePath)

	// 获取hallticket
	err = util.GetHallTicket(account)
	if err != nil {
		fmt.Printf("登录失败，原因:%s\n", err)
		return
	}

	// 获取jsessionid
	err = util.GetJSessionID()
	if err != nil {
		fmt.Printf("登录失败，原因:%s\n", err)
		return
	}

	// 获取浴室ID
	roomID := util.GetBathroomID()

	// 开始预订
	util.BookBathroom(roomID)

	// 为防止直接退出，进入输入等待状态
	_, _ = fmt.Scanln()

}

