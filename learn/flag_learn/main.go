/*
 * @Author: shenfuyuan
 * @Date: 2023-09-16 10:20:35
 * @LastEditTime: 2023-09-16 10:32:36
 * @Description:
 */
package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//定义命令行参数方式1
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	//解析命令行参数
	flag.Parse()

	fmt.Println(name, age, married, delay)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}

/*
.\main.exe -h
Usage of C:\Users\BIT_0306\Desktop\reddit\learn\flag_learn\main.exe:
  -age int
        年龄 (default 18)
  -d duration
        延迟的时间间隔
  -married
        婚否
  -name string
        姓名 (default "张三")

.\main.exe -name 申馥源 --age=23
申馥源 23 false 0s
[]
0
2 */
