package main

import (
	"fmt"
	"os"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(dir)
	inErr := creatFile("itibd.txt", " to day our class last day")
	fmt.Println(inErr)
}

func creatFile(fileName, content string) bool {
	posf, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer posf.Close()
	_, err = posf.Write([]byte(content))
	//fmt.Println(n, err)
	if err != nil {
		return false
	}
	return true
}
