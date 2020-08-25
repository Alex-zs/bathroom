package util

import (
	"fmt"
	"net/http"
	"strings"
)

func GetJSessionID() error {
	url := "https://vcard.ouc.edu.cn/Page/Page"
	method := "POST"

	payload := strings.NewReader("flowID=310&type=3&apptype=4&Url=https%253a%252f%252fvcard.ouc.edu.cn%253a20092%252fvenue.html%253ftypeid%253d1&MenuName=%u6D74%u5BA4%u9884%u8BA2&EMenuName=%u6D74%u5BA4%u9884%u8BA2&parm11=&parm22=&comeapp=1&headnames=&freamenames=&shownamess=&merpagess=&webheadhide=")

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Origin", "https://vcard.ouc.edu.cn")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Referer", "https://vcard.ouc.edu.cn/Phone/Index")
	req.Header.Add("Accept-Language", "en,zh-CN;q=0.9,zh;q=0.8")
	req.Header.Set("Cookie", fmt.Sprintf("ASP.NET_SessionId=%s; hallticket=%s;", AspNetSession, HallTicket))
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	resp, err := httpClient.Do(req)
	for _, cookie := range resp.Cookies() {
		if strings.HasPrefix(cookie.Name, "JSESSIONID") {
			JSessionID = cookie.Value
			return nil
		}
	}

	return nil
}
