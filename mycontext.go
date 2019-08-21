package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

type mCtx struct {
	start time.Time
	end time.Time
	done chan struct{}
}

func (c *mCtx) Deadline() (deadline time.Time, ok bool) {
	//debug.PrintStack()
	deadline = c.end
	ok = false
	return
}

func (c *mCtx) Done() <-chan struct{} {
	//debug.PrintStack()
	fmt.Println("self done", c.start, c.end)
	return c.done
}

func (c *mCtx) Err() error {
	return nil
}

func (c *mCtx) Value(key interface{}) interface{} {
	return nil
}

func sub(ctx context.Context) {
	fmt.Println(",,,,,,,,")
	select {
	case <- ctx.Done():
		fmt.Println("sub routine end")
	}
	fmt.Println(".........")
}

func main() {
	test()
	//ctx, cancelFunc := context.WithTimeout(context.Background(), 100 * time.Millisecond)
	//ctx, _ := context.WithTimeout(context.Background(), 10000 * time.Millisecond)
	ctx, cancelFunc := context.WithTimeout(&mCtx{time.Now(), time.Now().Add(1*time.Second), make(chan struct{})}, 2 * time.Second)
	fmt.Println(reflect.TypeOf(ctx))

	//cancelFunc()
	defer cancelFunc()

	go sub(ctx)
	select {
	case <-time.After(5*time.Second):
		fmt.Println("incase timeout")
	case <-ctx.Done():
		fmt.Println("context timeout", time.Now())
	}
	fmt.Println("main exit")
	time.Sleep(1*time.Second)
	fmt.Println("end")
}

type command interface {
	exec()
}

type obj1 struct {
	name string
}

func (o *obj1) exec() {
	fmt.Println(o.name)
}

type obj2 struct {
	i int
}

func (o *obj2) exec() {
	fmt.Println(o.i*3)
}

type bucket struct {
	cmds []command
}

func (bt *bucket) gen(cmd command) {
	fmt.Println("add", cmd)
	bt.cmds = append(bt.cmds, cmd)
}

func (bt *bucket) exec() {
	if len(bt.cmds) == 0 {
		fmt.Println("noting")
		return
	}

	for _, item := range bt.cmds {
		item.exec()
	}
}

func test() {
	fmt.Println("test")
	bt := &bucket{}
	bt.gen(&obj1{"ji"})
	bt.gen(&obj2{4})
	bt.gen(&obj1{"cai"})
	bt.exec()
}