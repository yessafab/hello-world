# hello-world
A simple Go app on Docker

## Setup


```
git clone https://github.com/yessafab/hello-world && cd hello-world
go test -v
docker build . -t hello-world
docker run --rm -p 3000:3333 hello-world
curl http://localhost:3000/
```
