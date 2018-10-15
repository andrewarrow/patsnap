package main

import "fmt"
import "os"

func main() {
	fmt.Println("pat, snap!")
	//DD_API
	//DD_APP
	api_key := os.Getenv("DD_API")
	app_key := os.Getenv("DD_APP")
	fmt.Printf("\n%s %s\n", api_key, app_key)
}
