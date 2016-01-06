package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	if data, err := ioutil.ReadDir("/home/pompeu/go/src/github.com/slides/src"); err == nil {
		for _, v := range data {
			fmt.Println(v.Name())
		}
	} else {
		fmt.Println(err)
	}
}
