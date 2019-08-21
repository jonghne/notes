package main

import (
	"fmt"
	dis "awesomeProject/discovery"
	"log"
	"time"
)

func main() {

	serviceName := "s-test"
	serviceInfo := dis.ServiceInfo{IP:"192.168.10.10", Desc:"super function"}

	s, err := dis.NewService(serviceName, serviceInfo,[]string {
		"http://172.16.0.3:2379",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)

	fmt.Println("check key")
	try := s.Get("jicai")
	for k, v := range try {
		log.Println(k, v)
	}

	go func() {
		time.Sleep(time.Second*20)
		s.Stop()
	}()

	go func() {
		for i:=0; i<60; i++ {
			time.Sleep(time.Second)
			ret := s.Get("services/s-test")
			for k, v := range ret {
				log.Println(k, v)
			}
		}
	}()

	s.Start()
}
