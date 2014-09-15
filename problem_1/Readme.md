# Problem 1

### Usage

    cd go_intro_problems/problem_1
    go build

write a go implementation of the md5sum program (in short, given a filename it outputs an md5 hash of its content)

bonus points if you can feed it by piping into it (cat file.txt | ./md5sum) or feed it from standard input (./md5sum < file.txt)

    go run md5sum.go < sample.txt

more bonus points if it can print an md5sum of itself if given no parameters at all :)

    go run md5sum.go

### Tests

    go test
