yohkan: *.go
	gofmt -w *.go
	gofmt -w */*.go
	go build -o yohkan *.go

.PHONY: callvis
callvis:
	go-callvis -format=png -file callvis -group pkg,type -nostd ./

.PHONY: clean
clean:
	rm *.png *.gv
