build:
	go build -o gostar main.go star.go
build-prod:
	go build -o /usr/local/bin/star main.go star.go
view:
	@cat ~/.star
clean:
	rm -f ~/.star