
build:
	go build -o ./build/varC ./cmd/varC/main.go 

dev: build
	./build/varC --file ./examples/example.varc

clean:
	rm  ./build/varC

lexer_test: build
	go test -v ./test/lexer_test.go 

utils_test: build
	go test -v ./test/utils_test.go 

parser_test: build
	go test -v ./test/parser_test.go 