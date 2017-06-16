package main

import (
	"net/http"
	"fmt"
	json2 "encoding/json"
)
type apistruct struct {
	Name string `json:"name"`
	Age string `json:"age"`
	Skill []string `json:"skill"`
}


func main() {
	http.HandleFunc("/api", apiGet)
	http.HandleFunc("/", firstPage)

	http.ListenAndServe(":8080", nil)

}

func firstPage(w http.ResponseWriter, r *http.Request)  {

	fmt.Fprint(w, "<h1 colro:red >Hello world")

}

func apiGet(w http.ResponseWriter, r *http.Request ) {

	data := apistruct{
		Name:"Best",
		Age:"23",
		Skill:[]string{"ios dev ", "golang", "andrion dev"},
	}

	json,err := json2.Marshal(data)
	if err != nil {
		panic(err)
	}

	//convert json
	api := string(json)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, api)
}