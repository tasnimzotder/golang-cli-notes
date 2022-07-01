.DEFAULT_GOAL := build

# variable
output_file := note

tidy:
	@go mod tidy

build: tidy
	@go build -o bin/$(output_file)

run: build
	@./bin/$(output_file)

clear:
	@rm -rf bin/*


