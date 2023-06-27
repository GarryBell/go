package main

import (
	"testing"
)

func TestResponseToString(t *testing.T) {
	var item Item
	item.Address.City = "Test"
	item.Address.Zipcode = "12345"
	item.Address.Geo.Lat = "123"
	item.Address.Geo.Lng = "456"

	got := responseToString(item)
	want := "Test 12345 (123, 456)"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
