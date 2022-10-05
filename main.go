package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {

	for j := 0; j < 5; j++ {

		go httpClient()

	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}

func httpClient() {
	url := "https://tls-test.eda-test.verification-gcp.colopl.jp/"

	for i := 0; i < 1; i++ {

		req, _ := http.NewRequest("GET", url, nil)

		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error Request:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			fmt.Println("Error Response:", resp.Status)
			return
		}

		//byteArray, _ := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(byteArray))
		fmt.Println(strconv.Itoa(i))
	}
}
