package main

import (
	"testing"
	"os"
	"bufio"
	"net/http"
	"fmt"
	"bytes"
	"strings"
	"encoding/json"
)

const filename = "./events.csv"
const url = "http://localhost:8080"

type CheckoutEvent struct{
	Count int `json:"checkout"`
}

func TestWritingToFile(t *testing.T) {
	initalLineCount := ReadLinesInFile(filename)
	_ = CreateNewEvent("checkout", "2015-02-11T15:01:00+00:00")
	finalLineCount := ReadLinesInFile(filename)
	if finalLineCount != initalLineCount + 1 {
		t.Error("Expected", initalLineCount + 1,  "got", finalLineCount)
	}
}

func TestStatusCodeReturned(t *testing.T) {
	resp := CreateNewEvent("checkout", "2015-02-11T15:01:00+00:00")
	if strings.Contains(resp.Status, "200") == false {
		t.Error("No response code of 200 returned, instead got", resp.Status)
	}
}

func TestJsonEncodingReturned(t *testing.T) {
	response := GetEvents("2015-02-11T15:01:00+00:00", "2015-02-11T15:01:00+00:00")
	if response.Header.Get("Content-Type") != "application/json"{
		t.Error("Content type application/json not found, instead got", response.Header.Get("Content-Type"))
	}
}

func TestAdditionalCheckoutEventEntered(t *testing.T) {
	
	response := GetEvents("2014-02-11T15:01:00+00:00", "2016-02-11T15:01:00+00:00")
	decoder := json.NewDecoder(response.Body)
	var checkoutEvent CheckoutEvent
	err := decoder.Decode(&checkoutEvent)
	if err != nil{
		fmt.Println("Error decoding event from response", err)
		return
	}
	initalCheckoutCount := checkoutEvent.Count

	_ = CreateNewEvent("checkout", "2015-02-11T15:01:00+00:00")

	response = GetEvents("2014-02-11T15:01:00+00:00", "2016-02-11T15:01:00+00:00")
	decoder = json.NewDecoder(response.Body)
	err = decoder.Decode(&checkoutEvent)
	if err != nil{
		fmt.Println("Error decoding event from response", err)
		return
	}
	finalCheckoutCount := checkoutEvent.Count
	if finalCheckoutCount !=  initalCheckoutCount+1 {
		t.Error("New checkout event not received", response.Status)
	}
}

func ReadLinesInFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil{
		fmt.Println("Cannot open file for reading", err)	
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan(){
		lineCount++
	}
	return lineCount
}

func CreateNewEvent(eventName string, eventTimestamp string)  *http.Response{
	var jsonStr = []byte(`{"name": "`+ eventName +`","timestamp": "`+ eventTimestamp +`"}`)
	req, err := http.NewRequest("POST", url+"/events", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp
}

func GetEvents(fromDate string, toDate string) *http.Response {
	response, err := http.Get(url+"/events/count?from="+fromDate+"&to="+toDate)
	if err !=nil{
		fmt.Printf("Error ", err)
		os.Exit(1)
	} 
	defer response.Body.Close()
	return response
}
