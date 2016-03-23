package main

import(	
	"encoding/json"
	"fmt"
	"net/http"
	"encoding/csv"
	"os"
	"time"
	"strings"
)

type Event struct{
	Name string `json:"name"`
	Timestamp string `json:"timestamp"`
}

const dateTimeLayout = "2006-01-02T15:04:05-07:00"
const filename = "./events.csv"

//Handle requests to /events
func recordEventsPost(rw http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)
	var event Event
	err := decoder.Decode(&event)
	if err != nil{
		http.Error(rw, "Error decoding event: " +err.Error(), 500)
	} else {
		err := writeEvent(event)
		if err != nil{
			http.Error(rw, "Error writing events to file: " +err.Error(), 500)
		}
	}
}

//write event to a file - could alternatively be implemented as a database write operation
func writeEvent(event Event) error{
	csvFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	eventString := []string{event.Name, event.Timestamp}
	err = writer.Write(eventString)
	writer.Flush()	
	return err
}

//Handle requests to /events/count
func aggregateEventsGet(rw http.ResponseWriter, request *http.Request){
	fromTime, err := getTimeFromQueryString(request.URL.RawQuery, "from")
	toTime, err := getTimeFromQueryString(request.URL.RawQuery, "to")
	allEvents, err := getAllEvents()
	eventNames := make(map[string]struct{})
	for _, event := range allEvents {
		eventNames[event.Name]=struct{}{}
	}
	eventMap := make(map[string]int)
	for eventName := range eventNames{
		eventCount := 0
		for _, event := range allEvents{
			eventTime, err := time.Parse(dateTimeLayout, event.Timestamp)
			if err != nil{
				http.Error(rw, "Error parsing event date: "+err.Error(), 500)
				return
			}
			if (eventName == event.Name && (eventTime.After(fromTime) && eventTime.Before(toTime))) {
				eventCount++
			}
		}
		eventMap[eventName]=eventCount
	}	
	jsonEventCount, err := json.Marshal(eventMap)
	if err != nil {
		http.Error(rw, "Error encoding JSON: "+err.Error(), 500)
	    return
	}
	fmt.Println(string(jsonEventCount))
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonEventCount)
}

//Using the date format specified, parse the time string
func getTimeFromQueryString(queryString string, parameterName string) (time.Time, error) {
	startIndex := (strings.Index(queryString, parameterName))+len(parameterName)+1
	endIndex := startIndex + len(dateTimeLayout)
	parsedTime, err := time.Parse(dateTimeLayout, queryString[startIndex:endIndex])
	if err != nil {
	    fmt.Println("Error parsing date string", err)
	}
	return parsedTime, err
}

//return a slice of all the events
func getAllEvents() ([]Event, error) {
	csvFile, err := os.Open(filename)
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	csvData, err := reader.ReadAll()
	var event Event
	var allEvents []Event
	for _, each := range csvData {
		event.Name = each[0]
		event.Timestamp = each[1]
		allEvents = append(allEvents, event)
	}
	return allEvents, err
}

func main(){
	fmt.Println("Events service started...")
	http.HandleFunc("/events/count", aggregateEventsGet)
	http.HandleFunc("/events", recordEventsPost)
	http.ListenAndServe(":8080", nil)
}