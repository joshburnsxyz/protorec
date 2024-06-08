build:
	go build \
		-v \
		-x \
		./cmd/protorec

clean:
	rm -f protorec

test:
	go test -v ./...
