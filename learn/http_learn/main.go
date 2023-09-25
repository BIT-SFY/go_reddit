package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//声明client 参数为默认
	client := &http.Client{}

	//声明要访问的url
	url := "http://www.baidu.com"

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)

	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)
	if err != nil {
		return
	}
	//返回的状态码
	status := response.Status
	fmt.Println()
	fmt.Println("1111111111111111111111", status)

}
