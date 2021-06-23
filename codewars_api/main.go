package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://www.codewars.com/api/v1/users/asgor"
  method := "GET"

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Cookie", "__cfduid=df4796c51279c66e6b7da967b9b19bbc11618668003")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  // receiving response
	respBody, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(respBody))

	// parsing response
	var cwData interface{}
	json.Unmarshal(respBody, &cwData) //extracting the json file
	//fmt.Println(cwData)

	// taking required values
	honor := cwData.(map[string]interface{})["honor"].(float64)
	ranks := cwData.(map[string]interface{})["ranks"]
	languages := ranks.(map[string]interface{})["languages"]
	golang := languages.(map[string]interface{})["go"]
	goScore := golang.(map[string]interface{})["score"].(float64)
	fmt.Println(honor, goScore)

	//return int(honor), int(goScore)
}