package client

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func calculateMD5(inputString string) string {
	h := md5.New()
	io.WriteString(h, inputString)
	encrypted := fmt.Sprintf("%x", h.Sum(nil))
	return encrypted
}

type BaiChuanClient struct {
	apiKey    string
	secretKey string
}

func NewBaiChuanClient(apiKey, secretKey string) *BaiChuanClient {
	return &BaiChuanClient{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (c *BaiChuanClient) ChatComplete(msg RequestMessage) (*ResponseMessage, error) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	timeStamp := strconv.Itoa(int(time.Now().Unix()))
	signature := calculateMD5(c.secretKey + string(jsonData) + timeStamp)

	url := "https://api.baichuan-ai.com/v1/chat"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("X-BC-Request-Id", msg.Id)
	req.Header.Set("X-BC-Timestamp", timeStamp)
	req.Header.Set("X-BC-Signature", signature)
	req.Header.Set("X-BC-Sign-Algo", "MD5")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		response := &ResponseMessage{}
		err = json.Unmarshal(body, response)
		if err != nil {
			return nil, err
		}
		return response, nil
	}
	return nil, fmt.Errorf("check status failed:%d", resp.Status)
}
