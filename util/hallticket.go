package util

import (
	"fmt"
	"strings"
)

func GetHallTicket(account string) error {
	resp, err := httpClient.Post("https://vcard.ouc.edu.cn/Login/QrCodeLogin",
		"application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("account=%s&json=true", account)))

	if err != nil {
		return err
	}

	for _, cookie := range resp.Cookies() {
		if strings.HasPrefix(cookie.Name, "hallticket") {
			HallTicket = cookie.Value
		}else if strings.HasPrefix(cookie.Name, "ASP") {
			AspNetSession = cookie.Value
		}
	}
	return nil
}
