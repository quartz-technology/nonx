BIN	=	nonx

all:
	go build -o $(BIN) main.go

clean:
	rm $(BIN)

re:	clean all

unit-tests:
	go test ./... -v -race

lint:
	golangci-lint run

.PHONY: all clean re unit-tests lint