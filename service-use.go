package main

import (
	"log"
	"time"
	"fmt"
	dis "awesomeProject/discovery"
)

func main() {

	m, err := dis.NewMaster([]string{
		"http://172.16.0.3:2379",
	}, "services/")

	if err != nil {
		log.Fatal(err)
	}

	for {
		for k, v := range  m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
		}
		fmt.Printf("nodes num = %d\n",len(m.Nodes))
		time.Sleep(time.Second * 5)
	}
}
