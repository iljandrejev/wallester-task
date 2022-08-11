package main

import "net/http"

func main() {
	err := http.ListenAndServe(":8080", ChiRouter().InitRouter())
	if err != nil {
		panic(err)
	}
}
