package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	for _, url := range urls {
		resp, err := http.Get(url)

		if err != nil {
			log.Println("cannot download CSV:", err)
			continue
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("cannot read CSV:", err)
			resp.Body.Close() // ensure the response body is closed
			continue
		}
		resp.Body.Close()
		ch <- b
	}
}

func main() {
	urls := []string{
		// 政府CIOポータルサイトのリンク
		"https://cio.go.jp/csv/BasicInformationAll-f2013.csv",
	}

	ch := make(chan []byte)

	var wg sync.WaitGroup
	wg.Add(1)

	go downloadCSV(&wg, urls, ch)


	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))

		for {
			records, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("error reading record:", err)
				break
			}
			// insertData(records)
			fmt.Println(records)
		}
	}
	wg.Wait()
}
