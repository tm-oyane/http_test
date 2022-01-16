package main

import (
	"fmt"
	"time"

	"net/http"
)

func main() {
	client := &http.Client{
		Timeout: HttpClientTimeout,
		//Transport: &http.Transport{
		//	DialContext: (&net.Dialer{
		//		KeepAlive: 15 * time.Second,
		//	}).DialContext,
		//	MaxIdleConnsPerHost: IdleConPerHost,
		//},
	}

	fmt.Println("Write csv")
	fmt.Println("start_unix_time,end_unix_time,latency_ms,req_cnt")
	for i := 1; i <= 1000; i++ {
		request(client, i)
		time.Sleep(time.Second * 1)
	}
}

const (
	CacheTtl = 86400 * time.Second // 1日
	HttpClientTimeout = 5 * time.Second
	IdleConPerHost = 10
	url = "http://localhost:8080/hello"
)

func request(client *http.Client, reqCnt int) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	start := timestamp()

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println(res.Header.Get("Connection"))
	//fmt.Println(res.Header.Get("Keep-Alive"))
	//fmt.Println(res.Header.Get("Content-Type"))

	//var data *Res
	//if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
	//}

	defer func() {
		//io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()

	end := timestamp()
	fmt.Println(fmt.Sprintf("%d,%d,%d,%d", start, end, end-start, reqCnt))
}

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

type Res struct {
	Message string `json:"message"`
}