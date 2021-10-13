package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func main() {
	data, err := ioutil.ReadFile("./images/16x16.pgm")	
	if err != nil {
		panic(err)
	}

	fields := strings.Fields(string(data))
	for i := range fields {
		fmt.Println(fields[i])
	}
}

