# Problem 3

### Usage

    cd go_intro_problems/problem_3

implement a mini clone of redis: a service with a REST API, that supports the following calls: 

* PUT - which records a key and a val; 
* GET - which given a key returns a val; 
* DELETE - which given a key deletes it as well as the value associated with it and 
* COUNT - which returns a count of all keys, optionally with a prefix of the key (e.g. if you have keys: "total_records", "total_bytes" "something_else", a COUNT "total" would return 2 while a COUNT with no param at all would return 3;

```
# start server
go run miniRedis.go

# PUT
curl -X PUT localhost:3000/records -d key1=val1
curl -X PUT localhost:3000/records -d key2=val2
curl -X PUT localhost:3000/records -d key3=val3
curl -X PUT localhost:3000/records -d foo1=bar1

# GET
curl localhost:3000/records?key=key1

# DELETE
curl -X "DELETE" localhost:3000/records?key=key1

# COUNT
curl localhost:3000/records/count
curl localhost:3000/records/count?q=key
```


### Tests

    ...
