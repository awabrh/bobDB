package main

import (
	"fmt"
	"math/rand"
	"os"
)

func SaveData(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, rand.Int())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	} 
	
	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}
	
	fp.Close()
	
	return os.Rename(tmp, path)
}

func main() {
	data := "Hello?!"
	err := SaveData("./output/text.txt", []byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}
}