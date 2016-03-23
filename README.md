<<<<<<< HEAD
# events_service

Events as a Service

Please provide a simple and well tested Go service that provides these pieces of functionality: 

● Send an event in the form of an HTTP POST that takes a name, and a timestamp 

● Query the events in a time range returning the name of the event, and the count of events in the time range 

An example API Recording an event

POST /events

{

name: 'checkout',

timestamp: '2015-02-11T15:01:00+00:00'

}

Aggregating events

GET /events/count?from=2015-02-11T15:01:00+00:00&to=2015-02-11T15:01:00+00:00

{

'checkout': 30

}

Other notes

● Dates should be ISO 8601 formatted date and time in UTC 

● Provide an explanation of the failure points, and why it will failure at scale 

● We like well tested code 

● Bonus points if it runs in Docker 
=======
# Events as a Service

I have created an events service that allows the user to:

* Send an event in the form of an HTTP POST that takes a name, and a timestamp 

* Query the events in a time range returning the name of the event, and the count of events in the time range 

An example API Recording an event:

###POST /events
```
{
"name": "checkout",
"timestamp": "2015-02-11T15:01:00+00:00"
}
```

###Aggregating events

```
GET /events/count?from=2015-02-11T15:01:00+00:00&to=2015-02-11T15:01:00+00:00
{
"checkout": 30
}
```

N.B. The request *MUST* be submitted in valid json format.

Not sure if your request is in valid Json format? You can check it here: http://jsonlint.com/


## Running the Events Service

In order to build the events service, run the following command:

`go build events_service.go`

In order to run the events service, run the following command after you have built the service:

`./events_service`


## Testing the Events Service

In order to test the events service, run the following command:

`go test events_service_test.go`
>>>>>>> f1b6198c9516d96a24994cc4cf6d16f318f97781
