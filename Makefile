run: build

build: main.go
	templ generate ./pages; go build -o ./bin
