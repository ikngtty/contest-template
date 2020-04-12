## How to run

```fish
# use `input.txt`
docker run -i --rm -v (pwd):/go/src golang:1.14.1 go run src/main.go < input.txt
```

or

```fish
# input stdio manually
docker run -it --rm -v (pwd):/go/src golang:1.14.1 go run src/main.go
```
