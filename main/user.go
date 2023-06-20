package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
)

var ids []int
var result Response

func main() {
	ids, result = loadIds()
	http.HandleFunc("/users/", usersHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func loadIds() ([]int, Response) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result Response
	var dummyData = []byte(`{"data":`)
	dummyData = append(dummyData, responseData...)
	dummyData = append(dummyData, '}')
	if err := json.Unmarshal(dummyData, &result); err != nil {
		fmt.Println("Cannot unmarshal JSON")
		fmt.Println(err)
	}
	var ids []int
	for i := 0; i < len(result.Data); i++ {
		ids = append(ids, result.Data[i].Id)
	}
	return ids, result
}

func handleInput(name string, w http.ResponseWriter) {
	if itemExists(ids, name) {
		var intVal, err = strconv.Atoi(name)
		if err == nil {
			for _, v := range result.Data {
				if v.Id == intVal {
					var item = v
					responseToString(item)
					var payload ReturnItem
					payload.Id = item.Id
					payload.Address = responseToString(item)
					res, _ := json.Marshal(payload)
					fmt.Fprintf(w, string(res))
				}
			}
		} else {

			fmt.Fprintf(w, "Error parsing id")
			fmt.Fprintf(w, string(err.Error()))
		}

	} else {
		fmt.Fprintf(w, "Name not found")
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	var name string = path.Base(r.URL.Path)
	handleInput(name, w)
}

func responseToString(item Item) string {
	return item.Address.City + " " + item.Address.Zipcode + " (" + item.Address.Geo.Lat + ", " + item.Address.Geo.Lng + ")"
}

func itemExists(arr []int, item string) bool {
	for i := 0; i < len(arr); i++ {
		if strconv.Itoa(arr[i]) == item {
			return true
		}
	}
	return false
}
