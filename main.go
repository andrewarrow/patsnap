package main

import "fmt"

//import "os"

import "encoding/json"

func main() {
	//fmt.Println("pat, snap!")
	//DD_API
	//DD_APP
	//	api_key := os.Getenv("DD_API")
	//app_key := os.Getenv("DD_APP")
	//fmt.Printf("\n%s %s\n", api_key, app_key)

	//https://app.datadoghq.com/api/v1/logs-queries/scopes/7156/list-csv?query=
	fmt.Println(makeJson())
}

func makeJson() string {
	time := map[string]interface{}{"offset": -7200000, "from": "now-604800s", "to": "now"}
	search := map[string]interface{}{"query": "@thing:866 AND @otherthing:foo"}
	f1 := map[string]interface{}{"field": map[string]interface{}{"path": "@bar"}}
	f2 := map[string]interface{}{"field": map[string]interface{}{"path": "@msg"}}
	columns := []interface{}{f1, f2}
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
