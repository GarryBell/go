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

func main() {
	http.HandleFunc("/users/", usersHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func processExternalAPICall(w http.ResponseWriter, response *http.Response) Response {
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result Response
	var dummyData = []byte(`{"data":`) // I'm sure there's some better way of doing this
	dummyData = append(dummyData, responseData...)
	dummyData = append(dummyData, '}')
	if err := json.Unmarshal(dummyData, &result); err != nil {
		fmt.Fprintf(w, "Cannot unmarshal JSON")
		fmt.Fprintf(w, err.Error())
	}
	var ids []int
	for i := 0; i < len(result.Data); i++ {
		ids = append(ids, result.Data[i].Id)
	}
	return result

}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	var name string = path.Base(r.URL.Path)
	var intVal, err = strconv.Atoi(name)
	if err != nil {
		fmt.Fprintf(w, "Error parsing id")
		return
	}

	// make api call
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal(err)
	}

	result := processExternalAPICall(w, response)

	for _, v := range result.Data {
		if v.Id == intVal {
			var item = v
			responseToString(item)
			var payload ReturnItem
			payload.Id = item.Id
			payload.Address = responseToString(item)
			res, _ := json.Marshal(payload)
			fmt.Fprintf(w, string(res))
			return
		}
	}
	fmt.Fprintf(w, "ID not found")
}

func responseToString(item Item) string {
	return item.Address.City + " " + item.Address.Zipcode + " (" + item.Address.Geo.Lat + ", " + item.Address.Geo.Lng + ")"
}
