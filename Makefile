clean:
	rm -rf slack

dev:
	go run main.go

build:
	go build -o slack main.go

build-arm:
	GOARCH=arm64 go build -o slack main.go

prod: build
	./slack

prod-test: clean
	@make build
	@./slack -w "" \
	-c chatbot-test -l "(1/3)" -r golang-slack-bot -b master -i ABCDE \
	-t leedonggyu -m hello-world -e dev -s heyboy  \
	-p https://google.com -o https://naver.com