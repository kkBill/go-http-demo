package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {

	// 1. 创建 Client
	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   500 * time.Second, // 指的是哪部分的超时时间？
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSClientConfig:       nil,
		TLSHandshakeTimeout:   10 * time.Second, // 指的是哪部分的超时时间？
		ResponseHeaderTimeout: 1 * time.Second,  // 指的是哪部分的超时时间？
	}

	client := &http.Client{
		Transport: trans,
		//CheckRedirect: nil, // 暂时不考虑
		//Jar:           nil, // 暂时不考虑
		Timeout: 10 * time.Second, // 指的是哪部分的超时时间？
	}

	// 2. 构造请求
	//volumeId := "1001"
	url := "http://localhost:8080/volume"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("new request error: %v\n", err)
		return
	}

	var resp *http.Response

	for i := 0; i < 10; i++ {
		// 3. 发起请求
		resp, err = client.Do(req)
		if err != nil {
			fmt.Printf("do request error: %v\n", err)
			return
		}

		fmt.Printf("Response status: %v\n", resp.Status)

		// 4. 读取数据并解析返回结果
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("read data from resp body error: %v\n", err)
			return
		}
		var volume Volume
		err = json.Unmarshal(bytes, &volume)
		if err != nil {
			fmt.Printf("json unmarshal error: %v\n", err)
			return
		}

		// 5. 处理返回结果
		fmt.Printf("Volume: %v\n", volume)
	}
	defer resp.Body.Close()
}
