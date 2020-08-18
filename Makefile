gelato: *.go
	gofmt -w *.go
	go build -o yohkan *.go
