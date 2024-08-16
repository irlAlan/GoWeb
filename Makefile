run: build
	./bin/main

build: main.go
	./bin/templ generate; go build -o ./bin main.go

clean:
	rm -rf pages/*_templ.go pages/layout/*_templ.go pages/blog/*_templ.go pages/blog/stem/*_templ.go
