GOX := $(shell which go)
BIN := protorec
SRC := ./cmd/protorec
OUT := ./dist

build:
	mkdir -p $(OUT)
	$(GOX) build \
		-v \
		-x \
		-o $(OUT)/$(BIN) \
		$(SRC)

clean:
	rm -f $(BIN)

test:
	go test -v ./...
