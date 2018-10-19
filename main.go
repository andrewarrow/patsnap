package main

import "fmt"
import "net/url"

import "strings"
import "os"
import "time"

func main() {
	cols := os.Args[1]
	tokens := strings.Split(cols, ",")
	acols := []string{}
	for _, t := range tokens {
		acols = append(acols, "\\\""+t+"\\\"")
	}
	collist := "\"[" + strings.Join(acols, ",") + "]\""
	fmt.Println(collist)
	query := url.QueryEscape(os.Args[2])
	from := time.Now().Unix() * 1000
	from = from - 15*60*1000
	to := time.Now().Unix() * 1000

	u := fmt.Sprintf("https://app.datadoghq.com/logs?cols=%s&event&from_ts=%d&index=main*live=true&query=%s&stream_sort=desc&to_ts=%d", url.QueryEscape(cols), from, query, to)
	fmt.Println(u)
	/*


		https://app.datadoghq.com/logs?cols=%22%5B%5C%22log_time%5C%22%2C%5C%22log_msg%5C%22%5D%22&event&from_ts=1539804802435&index=main&live=true&query=%40imei%3A866425031681804+%40dir%3Atb+%40msg%3A%2ARTO%2A&saved_view=5476&stream_sort=desc&to_ts=1539977602435

			https://app.datadoghq.com/logs?cols=%22%5B%5C%22log_time%5C%22%2C%5C%22log_msg%5C%22%5D%22&event&from_ts=1539804802435&index=main&live=true&query=%40imei%3A866425031681804+%40dir%3Atb+%40msg%3A%2ARTO%2A


			"[\"log_time\",\"log_msg\"]"&event&from_ts=1539804802435&index=main&live=true&query=@imei:866425031681804 @dir:tb @msg:*RTO*


			&saved_view=5476
			&stream_sort=desc
			&to_ts=1539977602435

	*/
}
