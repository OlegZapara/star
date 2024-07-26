build-dev:
	go build -o gostar main.go star.go
build:
	go build -o /usr/local/bin/star main.go star.go
clean:
	rm -f ~/.star