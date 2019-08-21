package main

import (
	"awesomeProject/redis"
	"fmt"
	db_config "github.com/dipperin/go-ms-toolkit/db-config"
	"github.com/dipperin/go-ms-toolkit/json"
	"net/url"
	"path"
	"time"
)

func newRedis() *redis.Redigo {
	red := redis.MakeRedigo(db_config.GetAppConfig().RedisUrl)
	return red
}

type info struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Addr string `json:"addr"`
	Career string `json:"career"`
	Hobby string `json:"hobby"`
	Age int `json:"age"`
}

func main() {
	conn := newRedis()
	err := conn.SetStr("mykey", "qy")
	if err != nil {
		fmt.Println("set", err)
		return
	}

	username, err := conn.GetStr("mykey")
	if err != nil {
		fmt.Println("get", err)
		return
	}
	fmt.Println(username)

	conn.Invalidate("mykey")
	username, err = conn.GetStr("mykey")
	if err != nil {
		fmt.Println("get", err)
		return
	}
	fmt.Println(username)
	//err = conn.SetHash("jc1", "loc1", "nanjing", "loc2", "beijing", "loc3", "haha")
	//if err != nil {
	//	fmt.Println("sethash", err)
	//	return
	//}
	//
	//err = conn.SetHash("jc2", "loc1", "tianjin", "loc2", "suzhou", "loc3", "lanzhou")
	//if err != nil {
	//	fmt.Println("sethash", err)
	//	return
	//}

	//personA := info{2, "luyun", "central avenue", "doctor", "foot", 23}
	//key := fmt.Sprintf("person_%d", personA.Id)
	//err = conn.SetHashStruct(key, personA)
	//if err != nil {
	//	fmt.Println("sethash", err)
	//	return
	//}
	//
	//var resp info
	//err = conn.GetHashStruct("person_2", &resp)
	//if err != nil {
	//	fmt.Println("get", err)
	//	return
	//}
	//fmt.Println(resp)
	//
	//personB := info{3, "qy", "7b", "stu", "bask", 73}
	//key = fmt.Sprintf("person_%d%d", personA.Id, personB.Id)
	//err = conn.SetHashStruct(key, []info{personA, personB})
	//if err != nil {
	//	fmt.Println("sethash", err)
	//	return
	//}

	fmt.Println(conn.GetStr("Peso2Go:0"))

	//testlist()
	//test_path()
	//
	//test_json()
}

func test_path() {
	fmt.Println(path.Join("haha", "hehe","mao", "xian"))
	u, err := url.Parse("www.baidu.com/?qiu=foot")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u, u.String())
}

type a struct {
	Name string
}

type b struct {
	Addr string
}

type c struct {
	M a `json:"xingming"`
	D b `json:"dizhi"`
}

type d struct {
	Id int `json:"id"`
	Con string `json:"con"`
	T  time.Time `json:"t"`
}

type dd struct {
	Con string `json:"con"`
	T  time.Time `json:"t"`
}

func test_json() {
	ra := a{"jicai"}
	rb := b{"shenzhen"}
	raw := c{ra, rb}

	str := json.StringifyJson(raw)
	fmt.Println(str)

	var inv c
	err:= json.ParseJson(str, &inv)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(inv)

	rdd := []dd{{Con:"hahaha", T: time.Now()}, {Con:"hehehe", T: time.Now()}, {T: time.Now()}}
	str = json.StringifyJson(rdd)
	fmt.Println(str)
	var rd []d
	err = json.ParseJson(str, &rd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rd)
}

func testlist() {
	s := []int{1,2,3,4,5,6,7,8,9,10,11,12}

	fmt.Printf("%p\n", s)
	a := s[:4]

	b := s[7:]
	fmt.Println(s, a, b)
	fmt.Printf("%p %p %p\n", s, a, b)
	fmt.Printf("%p %p\n", &s[0], &s[7])


	p:= append(a,b...)

	fmt.Println(s, p)
	fmt.Printf("%p %p\n", s, p)

}
