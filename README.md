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
