build:
	@go build -o bin/appender main.go
	@chmod +x bin/appender
	@echo "Build Done"

