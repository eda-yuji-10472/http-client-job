package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {

	for j := 0; j < 10; j++ {

		go httpClient()

	}

	time.Sleep(time.Second * 10)
	//fmt.Println("Complete")

	return

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

	//fmt.Println("gorutine complete")
}
