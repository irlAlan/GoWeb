run: build
	./bin/main

build: main.go
	./bin/templ generate; go build -o ./bin main.go
