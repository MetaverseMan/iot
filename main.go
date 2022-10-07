package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func main() {
	TestMqtt()
}
func TestMqtt() {
	var (
		clientId = "hp"
		wg       sync.WaitGroup
	)
	client := NewClient(clientId)
	err := client.Connect()
	if err != nil {
		fmt.Println("----------------")
		fmt.Println(err)
	}

	wg.Add(1)
	go func() {
		err := client.Subscribe(func(c *Client, msg *Message) {
			fmt.Printf("接收到消息: %+v \n", msg)
			wg.Done()
		}, 1, "news")

		if err != nil {
			panic(err)
		}
	}()

	msg := &Message{
		ClientID: clientId,
		Type:     "text",
		Data:     "Hello hp",
		Time:     time.Now().Unix(),
	}
	data, _ := json.Marshal(msg)

	err = client.Publish("news", 1, false, data)
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
