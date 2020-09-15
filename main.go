package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {
	username := flag.String("username", "", "username to check for")
	url := flag.String("urls", "https://2Dimensions.com/a/{}", "urls to check for")
	flag.Parse()

	var wg sync.WaitGroup

	urls := strings.Split(*url, ",")
	for _, u := range urls {
		url := strings.Replace(u, "{}", *username, 1)
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			resp, err := http.Head(url)
			if err != nil {
				log.Fatal(err)
			}

			if resp.StatusCode == http.StatusOK {
				fmt.Printf("[+] %s %s\n", url, http.StatusText(http.StatusOK))
			} else {
				fmt.Printf("[-] %s %s\n", url, http.StatusText(resp.StatusCode))
			}

		}(url)
	}
	wg.Wait()
}
