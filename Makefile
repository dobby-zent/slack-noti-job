clean:
	rm -rf slack

dev:
	go run main.go

build: clean
	go build -o slack main.go

build-arm: clean
	GOARCH=arm64 GOOS=linux go build -o slack main.go

prod: build
	./slack

prod-test: clean
	@make build
	@./slack -w "" \
	-c chatbot-test -l "(1/3)" -r golang-slack-bot -b master -i ABCDE \
	-t leedonggyu -m hello-world -e dev -s heyboy  \
	-p https://google.com -o https://naver.com

upload:
	aws s3 cp slack s3://zent-devops-jobs