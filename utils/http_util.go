package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	username = "cluster"
	password = "cluster"
)

func GetData(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	return string(b)
}
