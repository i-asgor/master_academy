package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://gmail.googleapis.com/gmail/v1/users/i.joni40@gmail.com/messages/1787936fa2a028c0"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer ya29.a0AfH6SMDYFYm55FoYg3bJO-d07u6Mualg8iDPgH-1LoW8xKFKcsSO8Gx27BAyVcSYmY672J1EIOK7rOEkFXP-7-CGa5BnafxtSBGaX7A53MC4Yh3EQ7iX3ZW6zE2au2w87cQEDe2t7z90ca2HO22odi1lRPsy")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(string(body))

}
