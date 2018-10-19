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
