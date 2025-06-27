package main

import (
	"fmt"
	router "myonly/internal/httpserver"
	"net/http"
)

func main() {
	mux := router.NewRouter()

	fmt.Println("서버 시작: http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
