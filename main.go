package main

import "fmt"
import "net/url"
import "io/ioutil"
import "net/http"

import "strings"
import "os"
import "time"

func main() {
	api_key := os.Getenv("DD_API")
	app_key := os.Getenv("DD_APP")
	/*
		    -d "api_key=${api_key}" \
				    -d "application_key=${app_key}" \
						    -d "from=${from_time}" \
								    -d "to=${to_time}" \
										    -d "query=system.cpu.idle{*}by{host}"
	*/
	from := 1541605519 - 3600
	to := 1541605519
	query := "system.cpu.idle{*}by{host}"
	url := fmt.Sprintf("https://api.datadoghq.com/api/v1/query?api_key=%s&application_key=%s&from=%d&to=%d&query=%s", api_key, app_key, from, to, query)
	d := DoGet(url)
	fmt.Println(d)
}
func DoGet(url string) string {
	request, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	resp, err := client.Do(request)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
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
func main2() {
	cols := os.Args[1]
	tokens := strings.Split(cols, ",")
	acols := []string{}
	for _, t := range tokens {
		acols = append(acols, "\"log_"+t+"\"")
	}
	collist := "[" + strings.Join(acols, ",") + "]"
	query := url.QueryEscape(os.Args[2])
	from := time.Now().Unix() * 1000
	from = from - 15*60*1000
	to := time.Now().Unix() * 1000

	u := fmt.Sprintf("https://app.datadoghq.com/logs?cols=%s&event&from_ts=%d&index=main*live=true&query=%s&stream_sort=desc&to_ts=%d", url.QueryEscape(collist), from, query, to)
	fmt.Println("")
	fmt.Println(u)
	fmt.Println("")
}
