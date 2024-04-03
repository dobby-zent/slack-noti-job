dev:
	go run main.go

build:
	go build -o slack main.go

prod: build
	./slack