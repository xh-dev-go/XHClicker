# Installation
```shell
go install github.com/xh-dev/XHClicker
```

# Issue when build (Not recommended)
In case problem when build in loading the bitmap.go file. 
Can add on head of bitmap.go :
```go
//go:build ignore
```