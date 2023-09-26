package main

import (
	"baichuan/client"
	"fmt"
)

func main() {
	apiKey := "1d878fd0c2672bf43005af9da6fd4520"
	secretKey := "I1GyxQcu5jFAPNJc0jxQO7GjQps="
	chatClient := client.NewBaiChuanClient(apiKey, secretKey)
	var request client.RequestMessage
	request.Id = "1"
	request.Model = "Baichuan2-53B"
	request.Parameter = client.ReqParameter{
		Temperature: 0.1,
	}
	request.Messages = []client.ReqMessage{
		{
			Role:    "user",
			Content: "hello world",
		},
	}

	response, err := chatClient.ChatComplete(request)
	if err != nil {
		fmt.Println(err)
	}

	if response.Code != 0 {
		fmt.Println("error:", response.Msg)
	} else {
		fmt.Println("success:", response.Data)
	}
}
