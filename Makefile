yohkan: *.go
	gofmt -w *.go
	go build -o yohkan *.go

callvis:
	go-callvis -format=png -file callvis -group pkg,type -nostd ./
