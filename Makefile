# application name
TARGET = stager_generator

build: clean
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./$(TARGET) main.go
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./$(TARGET).exe main.go

run:
	go run main.go

clean:
	rm -f ./$(TARGET)
	rm -f ./$(TARGET).exe