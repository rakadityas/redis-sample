run:
	@go build -o redis-sample && ./redis-sample

initialize:
	@go mod init | go mod vendor