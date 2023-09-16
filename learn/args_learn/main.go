package main

import (
	"fmt"
	"os"
)

// os.Args demo
func main() {
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}

/* go build .\main.go
.\main.exe 1 2 3 4
args[0]=C:\Users\BIT_0306\Desktop\reddit\learn\args_learn\main.exe
args[1]=1
args[2]=2
args[3]=3
args[4]=4 */
