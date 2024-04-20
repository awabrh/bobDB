package main

import (
	"log"
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

	err = fp.Sync()
	if err != nil {
		os.Remove(tmp)
		return err
	}
	
	fp.Close()
	
	return os.Rename(tmp, path)
}

func LogCreate(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0664)
}

func LogAppend(fp *os.File, line string) error {
	buf := []byte(line)
	buf = append(buf, '\n')
	_, err := fp.Write(buf)
	if err != nil {
		return err
	}

	return fp.Sync()
}

func main() {
	data := "Hello, World!"
	fp, err := LogCreate("./output/text.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = LogAppend(fp, data)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Success!")
	}
}