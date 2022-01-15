package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 27,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 10 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       10 * time.Second,
		},
	}

	fmt.Println("Write csv")
	fmt.Println("start_unix_time,end_unix_time,latency_ms,req_cnt")
	for i := 1; i <= 1000; i++ {
		request(client, i)
		time.Sleep(time.Second * 1)
	}
}

const url = "http://localhost:8080/hello"
func request(client *http.Client, reqCnt int) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	start := timestamp()

	res, err := client.Do(req)
	if err != nil {
		//fmt.Println("error req")
		return
	}

	//fmt.Println(res.Header.Get("Connection"))
	//fmt.Println(res.Header.Get("Keep-Alive"))
	//fmt.Println(res.Header.Get("Content-Type"))

	defer func() {
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()

	end := timestamp()
	fmt.Println(fmt.Sprintf("%d,%d,%d,%d", start, end, end-start, reqCnt))
}

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
