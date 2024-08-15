run: build

build: main.go
	./templ generate; go build -o ./bin
