# Common set of problems

1. write a go implementation of the md5sum program (in short, given a filename it outputs an md5 hash of its content); bonus points if you can feed it by piping into it (cat file.txt | ./md5sum) or feed it from standard input (./md5sum < file.txt); more bonus points if it can print an md5sum of itself if given no parameters at all :)
2. expose your md5sum as a service over HTTP: when POSTing a body it should respond with an md5 hash of the content;
3. implement a mini clone of redis: a service with a REST API, that supports the following calls: PUT - which records a key and a val; GET - which given a key returns a val; DELETE - which given a key deletes it as well as the value associated with it and COUNT - which returns a count of all keys, optionally with a prefix of the key (e.g. if you have keys: "total_records", "total_bytes" "something_else", a COUNT "total" would return 2 while a COUNT with no param at all would return 3;
4. implement matrix multiplication in Go: given an input file with lines in the form of "i j k", where i is row number, j is the column number and k is the value at row i, col j; then: if it is a square matrix, multiply it by its transpose otherwise return an error; use arrays for the matrices (and not slices);

Example output for #4:

    octave:1> A = [1 2 3; 2 2 2; -1 0 1]
    A =

       1   2   3
       2   2   2
      -1   0   1

    octave:2> A'
    ans =
    
       1   2  -1
       2   2   0
       3   2   1

    octave:3> A*A'
    ans =

       14   12    2
       12   12    0
        2    0    2

The #4 exercise is a bit contrived, the arrays may not be the best option for matrices (it depends on whether they are sparse or dense) but the point of the exercise is to get familiar with arrays vs. slices. So please try and find a way to use arrays even if it feels suboptimal. To simplify the problem, it should only handle matrices up to 10x10.

Use nothing but the standard library, you shouldn't need anything else to implement these. Try to keep things simple, always asking yourself "what's the simplest thing that would work?". That would be in the spirit of Go, and should lead you on the right path. And have fun! :)
