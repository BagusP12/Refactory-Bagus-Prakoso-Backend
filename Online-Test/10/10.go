package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Posts struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   bool   `json:"body"`
}

func main() {
	get()
}

func get() {
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var postsStruct Posts
	json.Unmarshal(bodyBytes, &postsStruct)
	fmt.Printf("API Response as struct:\n%+v\n", postsStruct)
}
