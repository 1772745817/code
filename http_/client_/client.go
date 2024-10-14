package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("Error ctreate newrequest", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error making new request ", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error read resp.body", err)
		return
	}

	fmt.Println("response: ", string(body))

}
