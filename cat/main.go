package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	size := len(os.Args)
	for i := 1; i < size; i++ {
		data, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(data))
	}
}
