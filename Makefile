GOX := $(shell which go)
BIN := protorec
SRC := ./cmd/protorec

build:
	$(GOX) build \
		-v \
		-x \
		-o $(BIN) \
		$(SRC)

clean:
	rm -f $(BIN)

test:
	go test -v ./...
