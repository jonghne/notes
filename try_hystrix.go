package main

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"time"
)

var Number int

var  Result string

type obj struct {
	num int
	ret string
}
func main() {
	config := hystrix.CommandConfig{
		Timeout:                2000, //超时时间设置  单位毫秒
		MaxConcurrentRequests:  10,    //最大请求数
		SleepWindow:            1,    //过多长时间，熔断器再次检测是否开启。单位毫秒
		ErrorPercentThreshold:  30,   //错误率
		RequestVolumeThreshold: 5,    //请求阈值  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
	}
	hystrix.ConfigureCommand("test", config)
	cbs, _, _ := hystrix.GetCircuit("test")
	defer hystrix.Flush()
	for i := 0; i < 50; i++ {
		start1 := time.Now()
		Number = i
		o := &obj{num:i}
		hystrix.Go("test", o.run1, o.getFallBack1)
		fmt.Println("请求次数:", i+1, ";用时:", time.Now().Sub(start1), ";请求状态 :", o.ret, ";熔断器开启状态:", cbs.IsOpen(), "请求是否允许：", cbs.AllowRequest())
		//hystrix.Go("test", run, getFallBack)
		//fmt.Println("请求次数:", i+1, ";用时:", time.Now().Sub(start1), ";请求状态 :", Result, ";熔断器开启状态:", cbs.IsOpen(), "请求是否允许：", cbs.AllowRequest())
		time.Sleep(1000 * time.Millisecond)
		//fmt.Println("请求次数:", i+1, ";用时:", time.Now().Sub(start1), ";请求状态 :", o.ret, ";熔断器开启状态:", cbs.IsOpen(), "请求是否允许：", cbs.AllowRequest())
	}
	time.Sleep(20 * time.Second)
}

func (o *obj) run1() error {
	o.ret = "RUNNING1"
	//fmt.Println(o.num+1, "in run")
	if o.num > 10 {
		return nil
	}

	if o.num%2 == 0 {
		return nil
	} else {
		//fmt.Println(o.num+1, "run fail")
		return errors.New("请求失败")
	}
}

func run() error {
	Result = "RUNNING1"
	if Number > 10 {
		return nil
	}
	if Number%2 == 0 {
		return nil
	} else {
		fmt.Println(Number, "run fail")
		return errors.New("请求失败")
	}
}

func (o *obj) getFallBack1(err error) error {
	o.ret = "FALLBACK"
	fmt.Println(o.num+1, "in fail", err)
	return nil
}

func getFallBack(err error) error {
	fmt.Println(err)
	Result = "FALLBACK"
	return nil
}
