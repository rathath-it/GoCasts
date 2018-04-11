# GoCasts

Companion repo to a course on Udemy.com


# Go to environment:

```
docker run -v $(pwd):/tmp -it golang:latest sh
```

# Run it inside Docker:

```sh
docker run -v $(pwd):/tmp -t golang:latest go run /tmp/code/school/main.go /tmp/code/school/model.go
```