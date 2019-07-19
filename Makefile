run: 
	cd cmd; \
	go run ./...

build:
	cd cmd; \
	go build -o ../bin/cprice ./...