package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	fileName := os.Getenv("FILE")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	buf := make([]byte, 32*1024) // define your buffer size here.

	for {
		n, err := file.Read(buf)

		if n > 0 {
			fmt.Print(string(buf[:n])) // your read buffer.
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read %d bytes: %v", n, err)
			break
		}
	}

	log.Println("Going to sleep")
	time.Sleep(10 * time.Minute)
	log.Println("Good morning")

	err = os.MkdirAll("/share/test/reports", os.ModePerm)
	if err != nil {
		panic(err)
	}
	d1 := []byte("hello\ngo\n")
	err = os.WriteFile("/share/test/reports/dat1", d1, 0755)
	if err != nil {
		panic(err)
	}
}
