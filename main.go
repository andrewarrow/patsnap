package main

import "fmt"
import "encoding/json"
import "io/ioutil"
import "net/http"

//import "compress/gzip"
import "os"
import "bytes"

func main() {
	from := 1541605519 - 3600
	to := 1541605519
	fmt.Println(from, to)
	d := DoPost()
	fmt.Println(d)
}

func DoPost() string {
	token := os.Getenv("DD_TOKEN")
	params := map[string]interface{}{}
	a2 := map[string]interface{}{}
	a3 := map[string]interface{}{}
	a4 := map[string]interface{}{}
	a4["query"] = "@port:8082"
	a3["offset"] = -7200000
	a3["from"] = "now-900s"
	a3["to"] = "now"
	a2["time"] = a3
	a2["search"] = a4
	params["aggregate"] = a2
	params["_authentication_token"] = token
	var buf, _ = json.Marshal(params)
	fmt.Println(string(buf))
	body := bytes.NewBuffer(buf)
	url := fmt.Sprintf("https://api.datadoghq.com/api/v1/logs-queries/scopes/main/aggregate")
	request, _ := http.NewRequest("POST", url, body)

	//request.Header.Add("Accept-Encoding", "gzip")
	request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{}

	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		//reader, err := gzip.NewReader(resp.Body)
		//body, err := ioutil.ReadAll(reader)
		if err == nil {
			fmt.Println(resp.StatusCode)
			if resp.StatusCode == 200 {
				return string(body)
			} else {
				fmt.Println(string(body))
			}
		} else {
			fmt.Println(string(body), err)
		}
	} else {
		fmt.Println(err)
	}
	return ""
}
