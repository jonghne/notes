package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

type myListener struct {
	*net.TCPListener
}

func newListener(addr string) (*myListener, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &myListener{ln.(*net.TCPListener)}, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	defer fmt.Fprintln(w, "hello world")

	fmt.Println("method:", r.Method)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	println("json:", string(body))

	var result []map[string]interface{}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("unmarshal fail: ", err)
	}
	fmt.Println(result)
}

func ParamHandler(w http.ResponseWriter, r *http.Request) {
	defer fmt.Fprintln(w, "hello world")
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.RawPath)
	fmt.Println(r.URL.Path, r.URL.String())
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.Header.Get("Token"))
}
func handleJack(w http.ResponseWriter, r *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		fmt.Fprintln(w, "fail")
	}
	conn, buf, _ := hj.Hijack()
	defer conn.Close()

	buf.WriteString("hi, hijack, money")

	buf.Flush()
}

//func main() {
//	http.HandleFunc("/standard", IndexHandler)
//	http.HandleFunc("/hi", handleJack)
//	http.ListenAndServe("127.0.0.1:8910", nil)
//}

func main() {
	ln, err := newListener("127.0.0.1:8910")
	if err != nil {
		fmt.Println("listener fail")
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/standard", http.HandlerFunc(IndexHandler))
	mux.Handle("/give", http.HandlerFunc(ParamHandler))
	mux.Handle("/hi", http.HandlerFunc(handleJack))
	(&http.Server{Handler:mux}).Serve(ln)
}
