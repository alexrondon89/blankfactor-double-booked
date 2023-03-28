## Double Booked

This is a simple api where a user can load his events for a day and
the application is going to detect what events are overlapped.

## Endpoints

This api is running over
```
http://127.0.0.1:3000/
```

two endpoints are available for this MVP
 - POST / addEvent - it adds a new event with a start and finish time.
```
Request:
{
    "event": [8.30, 10.30]
}

The start time is the position 0 in the array and the finish time is the position 1.
Is import to keep in mind that for this MVP the time is represented by a float where
the integer part is the hour and the decimals are the minutes.

Response: 
[8.3,10.3]

The response is going to be the same input array with an status response code 201
```

 - GET /collisions - it is going to return al the events that are being overlapping

``` 
Response: 
[
    {
        "baseEvent": [8.3, 10.3],
        "eventCollisions": [
            [8,9],
            [9,10],
            [10,11],
            [8.3,9.3]
        ]
    }
]

The response code is 200

In this case the event with start time 8:30am an finish time 10:30am
is being overlapped for all the events in the 'eventCollisions' array.
 
 i.e: the first item in eventCollisions array with start time 8am is finishing
 at 9am, it means, in the middle of baseEvent range time. This same logic apply for the
 rest of items in 'eventCollisions'
```

## Example

Adding all these events (one by one) using the endpoint POST /addEvent
```
[7,8]
[8, 9]
[9, 10]
[10, 11]
[8.30, 9.30]
[8.30, 10.30]
```

The response is going to be

```
[
    {
        "baseEvent": [7, 8],
        "eventCollisions": null
    },
    {
        "baseEvent": [10, 11],
        "eventCollisions": [
            [8.3, 10.3]
        ]
    },
    {
        "baseEvent": [8.3, 9.3],
        "eventCollisions": [
            [8, 9],
            [9, 10]
        ]
    },
    {
        "baseEvent": [8, 9],
        "eventCollisions": [
            [8.3, 9.3],
            [8.3, 10.3]
        ]
    },
    {
        "baseEvent": [9, 10],
        "eventCollisions": [
            [8.3, 9.3],
            [8.3, 10.3]
        ]
    },
    {
        "baseEvent": [8.3, 10.3],
        "eventCollisions": [
            [8, 9],
            [9, 10],
            [10, 11],
            [8.3, 9.3]
        ]
    }
]
```

## Considerations for this MVP
- I made use of goroutines and channel in validator logic to get a better performance. 
- The time is represented by a float where
the integer part is the hour and the decimals are the minutes.
- There is no validation around event input.
- The database is temporal, so if the server is down all the events saved are going to disappear

## TODO
- Add validation to input event to check for a valid hour and minutes.
- Do a more friendly POST /addEvent endpoint, allowing to the user add a event in a string way. 
For example: ["7:00", "16:00"] or with a json object.
- add logging track.
