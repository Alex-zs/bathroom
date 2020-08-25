package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var HallTicket, AspNetSession, JSessionID string


var BathroomName = []string{
	"鱼山校区-男浴室", "鱼山校区-女浴室",
	"北海苑-男浴室", "北海苑-女浴室",
	"东海苑-男浴室", "东海苑-女浴室",
	"南海苑-男浴室", "南海苑-女浴室",
	"浮山校区-男浴室", "浮山校区-女浴室",
}

var BathroomIDList = []string{
	"1526", "1525",
	"1544", "1543",
	"1546", "1545",
	"1548", "1547",
	"1566", "1565",
}

func GetBathroomID() string {
	roomIndex := -1

	fmt.Printf("浴室列表如下:\n")
	for i, s := range BathroomName {
		fmt.Printf("(%d): %s\n", i, s)
	}

	for true {
		fmt.Printf("请输入所要预订的浴室前面的数字并回车：")
		fmt.Scanln(&roomIndex)
		if roomIndex >= 0 && roomIndex < len(BathroomName) {
			fmt.Printf("你所预订的浴室为:%s\n", BathroomName[roomIndex])
			return BathroomIDList[roomIndex]
		}
	}
	return ""
}

func BookBathroom(roomID string) {
	date := time.Now().Format("2006-01-02")
	url := fmt.Sprintf("https://vcard.ouc.edu.cn:20092/product/findtime.html?type=day&s_dates=%s&serviceid=%s", date, roomID)

	attemptTime := 0
	for true {
		attemptTime++
		fmt.Printf("第%d次尝试中,请等待...\n", attemptTime)
		resp, err := httpClient.Get(url)
		if err != nil {
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)
		respJson := &CheckSupplyResp{}
		json.Unmarshal(body, respJson)

		if len(respJson.Object) == 1 && respJson.Object[0].SURPLUS > 0 {
			success := ToBook(respJson.Object[0].ID)
			if success {
				fmt.Printf("预订成功，请打开海大e卡通确认信息.")
				return
			}
		}
		time.Sleep(time.Second)
	}
}

func ToBook(id int) bool {
	url := "https://vcard.ouc.edu.cn:20092/order/tobook.html"
	method := "POST"

	payload := strings.NewReader("param=%7B%22stock%22%3A%7B%22" + strconv.Itoa(id) + "%22%3A%221%22%7D%2C%22extend%22%3A%7B%7D%7D&json=true")
	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("Origin", "https://vcard.ouc.edu.cn:20092")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Referer", "https://vcard.ouc.edu.cn:20092/product/show.html?id=1544")
	req.Header.Add("Accept-Language", "en,zh-CN;q=0.9,zh;q=0.8")
	req.Header.Set("Cookie", fmt.Sprintf("ASP.NET_SessionId=%s; JSESSIONID=%s; hallticket=%s", AspNetSession, JSessionID, HallTicket))

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("预订失败，原因：%s\n", err)
		return false
	}

	body, _ := ioutil.ReadAll(resp.Body)
	bookResp := &BookResp{}
	json.Unmarshal(body, bookResp)
	if bookResp.Message == "预订成功" {
		return true
	}else {
		fmt.Printf("预订失败，原因：%s\n", bookResp.Message)
		return false
	}
}


// 检查浴室剩余数量的相应
type CheckSupplyResp struct {
	Object []struct {
		PRICE     string `json:"PRICE"`
		SURPLUS   int    `json:"SURPLUS"`
		USINGNUM  int    `json:"USING_NUM"`
		SERVICEID string `json:"SERVICEID"`
		ID        int    `json:"ID"`
		REMARK    string `json:"REMARK"`
		TIMENO    string `json:"TIME_NO"`
		STATUS    int    `json:"STATUS"`
		ISOVER    int    `json:"ISOVER"`
		ALLCOUNT  int    `json:"ALL_COUNT"`
	} `json:"object"`
	Result string `json:"result"`
}

type BookResp struct {
	Message    string      `json:"message"`
}