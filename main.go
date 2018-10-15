package main

import "fmt"
import "strings"
import "os"
import "encoding/json"

func main() {
	//fmt.Println("pat, snap!")
	//DD_API
	//DD_APP
	//	api_key := os.Getenv("DD_API")
	//application_key := os.Getenv("DD_APP")
	//fmt.Printf("\n%s %s\n", api_key, app_key)

	//https://app.datadoghq.com/api/v1/logs-queries/scopes/7156/list-csv?query=
	query := os.Args[1]
	columns := os.Args[2]
	fmt.Println(makeJson(query, columns))
}

func makeJson(query, columnString string) string {
	time := map[string]interface{}{"offset": -7200000, "from": "now-604800s", "to": "now"}
	search := map[string]interface{}{"query": query}
	columns := []interface{}{}
	tokens := strings.Split(columnString, " ")
	for _, t := range tokens {
		f := map[string]interface{}{"field": map[string]interface{}{"path": t}}
		columns = append(columns, f)
	}
	t1 := map[string]interface{}{"time": map[string]interface{}{"order": "desc"}}
	sort := map[string]interface{}{"sort": t1}
	list := map[string]interface{}{"time": time, "search": search, "columns": columns, "limit": 5000, "sort": sort}
	payload := map[string]interface{}{
		"list": list}

	b, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return (string(b))
}
