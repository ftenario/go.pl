package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchurl(url, ch)
	}

	for _, v := range os.Args[1:] {
		//fmt.Println(<-ch) // receive data from channel
		dir, _ := os.Getwd()
		file := v[strings.Index(v, "/"):]
		fh, err := os.Create(path.Join(dir, file))
		if err != nil {
			log.Println("Error:", err)
		}
		defer fh.Close()
		str := <-ch
		_, err = fh.Write([]byte(str))
		if err != nil {
			log.Println("Error:", err)
		}
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchurl(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	ch <- buf.String()
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs %7d %s\n", secs, buf.Len(), url)
}
