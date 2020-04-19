build:
	protoc -I. --go_out=plugins=micro:. proto/demo/demo.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t laracom-demo-service .



