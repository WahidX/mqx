gin:
	gin --build cmd -i --bin build/server

test:
	go test -v ./...