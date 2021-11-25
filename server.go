package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 0,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       0,
	}

	http.HandleFunc("/volume", GetVolumeById)

	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server listen and serve error: %v\n", err)
		return
	}
}

func GetVolumeById(writer http.ResponseWriter, request *http.Request) {
	volume := Volume{
		Capacity: 10,
		Region:   "shanghai",
		VolumeID: "1001",
	}
	bytes, err := json.Marshal(&volume)
	if err != nil {
		fmt.Printf("json marshal error: %v\n", err)
		return
	}

	// 表示处理时间
	writer.WriteHeader(200)

	_, err = writer.Write(bytes)
	time.Sleep(3 * time.Second)
	if err != nil {
		fmt.Printf("write error: %v\n", err)
		return
	}
}
