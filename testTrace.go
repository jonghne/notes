package main

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
	"log"
	"io/ioutil"
)

func main() {
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)

	// 创建客户端请求跟踪
	trace := &httptrace.ClientTrace{
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			fmt.Println(info.Host)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println(network)
			fmt.Println(addr)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
	}
	// 请求追踪
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
	data,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}