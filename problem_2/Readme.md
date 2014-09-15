# Problem 2

### Usage

    cd go_intro_problems/problem_2

expose your md5sum as a service over HTTP: when POSTing a body it should respond with an md5 hash of the content;

    # start server
    go run md5service.go

    # send Post request
    curl -d 'Lorem ipsum dolor sit amet.' http://localhost:3000/md5

### Tests

    ...
