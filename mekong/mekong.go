package mekong

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// apiUrl Mekong 美控短信
const apiUrl = "https://api.mekongsms.com/api/postsms.aspx"

func md5Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func (m *Mekong) GetUserName() string {
	return m.UserName
}

func (m *Mekong) GetPassword() string {
	return md5Hash(m.Password)
}

func (m *Mekong) SendSMS(sender string, content string, gsm string, i int) (responseContent string, err error) {
	requestData := map[string]string{
		"username": m.GetUserName(),
		"pass":     m.GetPassword(),
		"sender":   sender,
		"smstext":  content,
		"gsm":      gsm,
		"int":      strconv.Itoa(i),
		"cd":       strconv.FormatInt(time.Now().UnixNano(), 10),
	}

	formData := url.Values{}
	for key, value := range requestData {
		formData.Set(key, value)
	}

	resp, err := http.PostForm(apiUrl, formData)
	if err != nil {
		fmt.Println("form data post err: ", err)
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("form data post body: ", string(body))
	return string(body), nil
}
