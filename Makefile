build:
	go build \
		-v \
		-x \
		./cmd/protorec

test:
	go test -v ./...
