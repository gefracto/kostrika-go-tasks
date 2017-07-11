package main

import (
	"net/http"

	"github.com/gefracto/kostrika-go-tasks/src/server"
	_ "github.com/gefracto/kostrika-go-tasks/src/task1"
	_ "github.com/gefracto/kostrika-go-tasks/src/task2"
	_ "github.com/gefracto/kostrika-go-tasks/src/task3"
	_ "github.com/gefracto/kostrika-go-tasks/src/task4"
	_ "github.com/gefracto/kostrika-go-tasks/src/task5"
	_ "github.com/gefracto/kostrika-go-tasks/src/task6"
	_ "github.com/gefracto/kostrika-go-tasks/src/task7"
)

func main() {

	http.HandleFunc("/task/", server.HandleTask)
	http.ListenAndServe(":1111", nil)
}
