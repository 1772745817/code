package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	jsondata := `{"title": "foo", "body": "bar", "userId": 1}`

	body := bytes.NewBuffer([]byte(jsondata))

	resp, err := client.Post("https://jsonplaceholder.typicode.com/posts", "application.json", body)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	fmt.Println("resp: ", resp.Status)

}
