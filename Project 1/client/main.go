package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createBook() int64 {
	url := "http://localhost:3000/api/v1/book"
	payload := []byte(`{"name": "Sample Book", "author": "Sample Author", "rating": 5}`)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		var response map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			panic(err)
		}
		id, ok := response["ID"].(float64)
		if !ok {
			panic("Failed to parse ID from response")
		}
		return int64(id)
	} else {
		panic("Failed to create book")
	}
}

func getBooks() {
	url := "http://localhost:3000/api/v1/book"

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func getBookById(id int64) {
	url := "http://localhost:3000/api/v1/book/" + fmt.Sprint(id)

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func deleteBookById(id int64) {
	url := "http://localhost:3000/api/v1/book/" + fmt.Sprint(id)

	req, _ := http.NewRequest("DELETE", url, nil)

	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		panic(err)
	}
}

func main() {
	createdBookID := createBook()
	getBooks()
	getBookById(createdBookID)
	deleteBookById(createdBookID)
}
