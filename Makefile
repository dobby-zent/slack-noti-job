## Not Call
clean:
	rm -rf slack-*

## Not Call
build-linux:
	go build -o slack-linux main.go

## Not Call
build-arm: 
	GOARCH=arm64 GOOS=linux go build -o slack-arm64 main.go

######################################## Use this ########################################
build: clean
	@make build-linux
	@make build-arm

dev:
	go run main.go

prod:
	./slack-linux

prod-test: clean
	@make build
	@./slack -w "" \
	-c chatbot-test -l "(1/3)" -r golang-slack-bot -b master -i ABCDE \
	-t leedonggyu -m hello-world -e dev -s heyboy  \
	-p https://google.com -o https://naver.com

upload: build
	aws s3 cp slack-arm64 s3://zent-devops-jobs/slack --profile zent-dev
	aws s3 cp slack-arm64 s3://zent-devops-shared-jobs/slack --profile zent-root