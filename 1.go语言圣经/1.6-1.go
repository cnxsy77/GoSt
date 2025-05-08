// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
1.6 并发获取多个URL
*/
func main() {
	statr := time.Now()
	ch := make(chan string)
	for num, url := range os.Args[1:] {
		go fetch(url, ch, num)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)

	}
	fmt.Printf("%.2fs elapsed\n", time.Since(statr).Seconds())

}

func fetch(url string, ch chan<- string, num int) {
	statr := time.Now()
	var utl_fime_name string
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nums := strconv.Itoa(num)
	utl_fime_name = nums + ".txt"
	dst, err := os.Create(utl_fime_name)
	if err != nil {
		panic(err)
	}
	nbytes, err := io.Copy(dst, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(statr).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
