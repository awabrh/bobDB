package main

import (
	"fmt"
	"os"
)

func SaveData(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	} 

	defer fp.Close()

	_, err = fp.Write(data)
	return err
}

func main() {
	data := "Hello, again!"
	SaveData("./out/test.txt", []byte(data))
	fmt.Println(data)
}