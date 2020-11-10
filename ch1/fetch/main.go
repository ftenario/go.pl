package main

import (
	"fmt"
	"go/src/bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// 1.8 - if no http://, add it
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Error: %+v", err)
		}

		//using ioutil
		// b, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Printf("Error: %+v", err)
		// }

		//1.7 - Using io.Copy
		buf := bufio.NewWriter(os.Stdout)
		_, err = io.Copy(buf, resp.Body)
		if err != nil {
			log.Printf("%+v", err)
		}

		defer resp.Body.Close()

		fmt.Printf("%s\n", buf)

		//1.9 - Print HTTP status
		fmt.Println("Http Status:", resp.Status)
	}
}
