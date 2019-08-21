package main

import (
	"errors"
	"fmt"
)

type executor struct {
	exception   chan interface{}
	defaultFunc func(err error)
}

func New(f func(err error)) *executor {
	ech := make(chan interface{})
	return &executor{exception: ech, defaultFunc: f}
}

func (exe *executor) try(f func()) *executor {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				exe.exception <- err
			} else {
				exe.exception <- nil
			}
		}()

		f()
	}()

	return exe
}

func (exe *executor) catch(f func(err error)) {
	//wait result
	ret := <- exe.exception
	if ret == nil {
		fmt.Println("exec ok")
		return
	}
	if f != nil {
		f(ret.(error))
	} else {
		exe.defaultFunc(ret.(error))
	}
}

func proc1(err error) {
	fmt.Println("default process", err)
}

func proc2(err error) {
	fmt.Println("self process", err)
}

func retVal() int {
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	panic("what")
	return 34
}

func main() {
	fmt.Println(retVal())

	exe := New(proc1)

	exe.try(func() {
		fmt.Println("maoxian")
		panic(errors.New("boom"))
	}).catch(nil)

	exe.try(func() {
		fmt.Println("success")
	}).catch(proc2)
}
