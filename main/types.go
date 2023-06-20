package main

type Response struct {
	Data []Item
}

type Item struct {
	Id      int `json:"id"`
	Address struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
}

type ReturnItem struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
}
