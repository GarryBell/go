package main

import "testing"

func TestItemExists(t *testing.T) {
	arr := []int{1, 2, 3}

	got := itemExists(arr, "2")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	got = itemExists(arr, "4")
	want = false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

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
