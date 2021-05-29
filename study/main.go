package main

import (
	"errors"
	"fmt"
)

type T struct {

}

func main()  {
	err := errors.New("测试错误")
	CheckErr(err)
	fmt.Println("123")
}

func CheckErr(err error)  {
	if err != nil{
		Try(func() {
			panic("测试错误")
		}, func(err interface{}) {
			fmt.Println(err)
			return
		})
	}
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
